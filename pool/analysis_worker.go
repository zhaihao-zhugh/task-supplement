package pool

import (
	"encoding/json"
	"fmt"
	h "gpk/http"
	"gpk/logger"
	"io"
	"supplementary-inspection/model"

	"github.com/google/uuid"
)

var AnalysisHost model.Host
var HttpHost model.Host
var AnalyzeTimeout int

type AnalysisWorker struct {
	ErrorChan chan error
	Wc        chan *model.AnalysisResult
	RequestID string // 请求ID
}

func NewAnalysisWorker() *AnalysisWorker {
	return &AnalysisWorker{
		ErrorChan: make(chan error),
		Wc:        make(chan *model.AnalysisResult),
		RequestID: uuid.New().String(),
	}
}

// Work 发送分析请求
func (worker *AnalysisWorker) Work(items []model.AnalysisItem) {
	logger.Infof("开始等待分析主机返回分析结果,超过%d秒后算超时", AnalyzeTimeout)

	// res := make(chan error)

	// go func() {
	// 	select {
	// 	// 分析过程超时
	// 	case <-time.After(time.Second * time.Duration(AnalyzeTimeout)):
	// 		// for _, item := range items {
	// 		// 	item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析过程超时")
	// 		// 	logger.Errorf("###分析过程超时:%s", item.Point.Name)
	// 		// }
	// 		logger.Errorf("请求识别唯一标识为%s的图像分析过程超时", worker.RequestID)
	// 		res <- fmt.Errorf("分析主机 图像分析超时")
	// 		return
	// 	// 处理分析结果
	// 	case result := <-worker.Wc:
	// 		logger.Infof("正在处理分析结果:%s", worker.RequestID)
	// 		for _, item := range items {
	// 			// exist := false
	// 			for _, object := range result.ResultsList {

	// 				// if object.ObjectID == item.ObjectID {
	// 				// 	exist = true
	// 				// 	// item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(&object, &item, "")
	// 				// }

	// 				if object.ResImageBase64 != "" {
	// 					pic_data, err := service.CovertBase64ToPic(object.ResImageBase64)
	// 					if err != nil {
	// 						logger.Errorf("base64图片解析错误:%s", err.Error())
	// 					} else {
	// 						logger.Info("保存识别图片")
	// 						service.SaveBytesToFile(pic_data, "./"+item.Point.Name+"_det.jpg")
	// 					}
	// 				}
	// 			}
	// 			// if !exist {
	// 			// 	// item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析主机漏检")
	// 			// 	logger.Errorf("###分析主机漏检:%s", item.Point.Name)
	// 			// }
	// 		}

	// 		res <- nil
	// 		logger.Info("分析结果处理完成")
	// 		return
	// 	}
	// }()

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

	// 网络代理
	// buf, err := h.Post("https://221.226.190.26:18443/detect/picAnalyse", &request)

	if err != nil {
		for _, item := range items {
			// item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "请求发送错误")
			logger.Errorf("###请求发送错误:%s", item.Point.Name)
		}
		logger.Errorf("请求识别唯一标识为%s的图像分析请求发送错误%s", worker.RequestID, err.Error())
		worker.ErrorChan <- err
	} else {
		logger.Infof("请求识别唯一标识为%s的图像分析请求发送成功%v", worker.RequestID, buf)
	}
	body, _ := io.ReadAll(buf.Body)
	defer buf.Body.Close()
	data := make(map[string]int)
	_ = json.Unmarshal(body, &data)

	if code := data["code"]; code != 200 {
		for _, item := range items {
			// item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析返回错误")
			logger.Errorf("###分析返回错误:%s", item.Point.Name)
		}
		logger.Errorf("请求识别唯一标识为%s的图像分析请求返回错误", worker.RequestID)
		worker.ErrorChan <- fmt.Errorf("图像分析错误")
	}
}
