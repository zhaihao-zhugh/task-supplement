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

	logger.Infof("接收图像分析结果,原始数据为:%s", string(buf))

	content := gin.H{}

	result := new(model.AnalysisResult)

	if err := json.Unmarshal(buf, &result); err == nil {

		exist := false

		pool.PAnalysisRunner.Workers.Range(func(key, value interface{}) bool {
			worker := value.(*pool.AnalysisWorker)
			if result.RequestID == worker.RequestID {
				worker.Wc <- result
				exist = true
				content["code"] = 200
				logger.Infof("接收到请求识别唯一标识为%s的图像分析结果", result.RequestID)
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
