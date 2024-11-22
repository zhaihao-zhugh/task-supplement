package api

import (
	"encoding/json"
	"gpk/logger"
	"io"
	"net/http"
	"supplementary-inspection/model"
	"supplementary-inspection/pool"

	"github.com/gin-gonic/gin"
)

func PicAnalyseRetNotify(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
		}
	}()

	buf, _ := io.ReadAll(ctx.Request.Body)

	logger.Infof("接收图像分析结果")

	content := gin.H{}

	result := new(model.AnalysisResult)

	if err := json.Unmarshal(buf, &result); err == nil {

		exist := false
		logger.Infof("分析结果唯一标识为%s", result.RequestID)
		pool.PAnalysisRunner.Workers.Range(func(key, value interface{}) bool {
			worker := value.(*pool.AnalysisWorker)
			logger.Infof("接收到请求识别唯一标识为%s", worker.RequestID)
			if result.RequestID == worker.RequestID {
				logger.Infof("接收到请求识别唯一标识为%s的图像分析结果", result.RequestID)
				worker.Wc <- result
				exist = true
				content["code"] = 200
				logger.Info("发送结果成功")
				return false
			}
			return true
		})

		if !exist {
			content["code"] = 400
			logger.Errorf("接收到请求识别唯一标识为%s的图像分析结果,但与请求标识不一致", result.RequestID)
		}
	} else {
		content["code"] = 400
		logger.Errorf("图像分析结果解析错误:%s", err.Error())
	}

	ctx.JSON(http.StatusOK, content)
}
