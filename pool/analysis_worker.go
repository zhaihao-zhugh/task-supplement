package pool

import (
	"encoding/json"
	"fmt"
	h "gpk/http"
	"gpk/logger"
	"io"
	"supplementary-inspection/model"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var AnalysisHost model.Host
var HttpHost model.Host
var AnalyzeTimeout = viper.GetInt("settings.analysis-timeout")

type AnalysisWorker struct {
	Wc        chan *model.AnalysisResult
	RequestID string // 请求ID
}

func NewAnalysisWorker() *AnalysisWorker {
	return &AnalysisWorker{
		Wc:        make(chan *model.AnalysisResult),
		RequestID: uuid.New().String(),
	}
}

// Work 发送分析请求
func (worker *AnalysisWorker) Work(ch chan struct{}, items []model.AnalysisItem) error {
	defer func() {
		<-ch
	}()

	logger.Infof("开始等待分析主机返回分析结果,超过%d秒后算超时", AnalyzeTimeout)

	res := make(chan error)

	go func() {
		select {
		// 分析过程超时
		case <-time.After(time.Second * time.Duration(AnalyzeTimeout)):
			for _, item := range items {
				item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析过程超时")
				logger.Errorf("###分析过程超时:%s", item.Point.Name)
			}
			logger.Errorf("请求识别唯一标识为%s的图像分析过程超时", worker.RequestID)
			res <- fmt.Errorf("分析主机 图像分析超时")
			return
		// 处理分析结果
		case result := <-worker.Wc:
			logger.Infof("正在处理分析结果:%s", worker.RequestID)
			for _, item := range items {
				exist := false
				for _, object := range result.ResultsList {
					if object.ObjectID == item.ObjectID {
						exist = true
						item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(&object, &item, "")
					}
				}
				if !exist {
					item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析主机漏检")
					logger.Errorf("###分析主机漏检:%s", item.Point.Name)
				}
			}
			res <- nil
			return
		}
	}()

	logger.Infof("发送请求识别唯一标识为%s的图像分析请求,结果集长度为%d", worker.RequestID, len(items))

	request := model.AnalysisRequest{
		RequestHostIP:   HttpHost.IP,
		RequestHostPort: fmt.Sprintf("%d", HttpHost.Port),
		RequestID:       worker.RequestID,
		ObjectList:      items,
	}

	buf, err := h.Post(fmt.Sprintf("http://%s:%d/picAnalyse",
		AnalysisHost.IP,
		AnalysisHost.Port,
	), &request)

	if err != nil {
		for _, item := range items {
			item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "请求发送错误")
			logger.Errorf("###请求发送错误:%s", item.Point.Name)
		}
		logger.Errorf("请求识别唯一标识为%s的图像分析请求发送错误%s", worker.RequestID, err.Error())
		return err
	} else {
		logger.Infof("请求识别唯一标识为%s的图像分析请求发送成功%v", worker.RequestID, buf)
	}
	body, _ := io.ReadAll(buf.Body)
	defer buf.Body.Close()
	data := make(map[string]int)
	_ = json.Unmarshal(body, &data)

	if code := data["code"]; code != 200 {
		for _, item := range items {
			item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析返回错误")
			logger.Errorf("###分析返回错误:%s", item.Point.Name)
		}
		logger.Errorf("请求识别唯一标识为%s的图像分析请求返回错误", worker.RequestID)
		return fmt.Errorf("图像分析错误")
	}

	return <-res
}
