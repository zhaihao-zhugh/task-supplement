package route

import (
	"gpk/http"
	"gpk/logger"
	"supplementary-inspection/api/v1"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	var router = gin.Default()
	BaseRouter := router.Group("")
	{
		BaseRouter.POST("GetTaskContent", api.GetTaskContent)
		BaseRouter.POST("UploadTaskResult", api.UploadTaskResult)
		BaseRouter.POST("picAnalyseRetNotify", api.PicAnalyseRetNotify)
	}
	return router
}

func RunHttpServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	http.SetHandler(Router())
	logger.Error(http.Run(port))
}
