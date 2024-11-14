package api

import (
	"encoding/json"
	"gpk/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadTaskResultRequest struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		MainTask struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			Name    string `json:"name"`
			SubTask []struct {
				Id        string `json:"id"`
				Type      string `json:"type"`
				Name      string `json:"name"`
				Clearance []struct {
					Id        string `json:"id"`
					Name      string `json:"name"`
					Sn        string `json:"sn"`
					TestPoint []struct {
						Id       string `json:"id"`
						Name     string `json:"name"`
						Sn       string `json:"sn"`
						FileName string `json:"filename"`
						Result   int    `json:"result"`
					} `json:"test_point"`
				} `json:"clearance"`
			} `json:"sub_task"`
		} `json:"main_task"`
	} `json:"data"`
}

func UploadTaskResult(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
		}
	}()

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取文件
	files := form.File["files"]
	for _, file := range files {
		// 处理文件，例如保存到指定位置
		logger.Infof("接收到文件:%+v\n", file)
		err := ctx.SaveUploadedFile(file, "/store/ftp/tmp/"+file.Filename)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// 获取结果数据
	data := form.Value["data"][0]
	logger.Infof("接收到任务结果上报:%+v\n", data)
	var request UploadTaskResultRequest
	err = json.Unmarshal([]byte(data), &request)
	if err != nil {
		logger.Errorf("%+v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	// for _, subTask := range request.Data.MainTask.SubTask {
	// 	for _, clearance := range subTask.Clearance {
	// 		for _, point := range clearance.TestPoint {

	// 			item := model.AnalysisItem{
	// 				ObjectID:   point.Id,
	// 				ObjectName: point.Name,
	// 			}
	// 		}
	// 	}
	// }

	ctx.JSON(http.StatusOK, request)
}
