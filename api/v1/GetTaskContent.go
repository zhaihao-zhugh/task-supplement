package api

import (
	"encoding/json"
	"gpk/logger"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type GetTaskResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		MainTask []struct {
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
						Id        string `json:"id"`
						Name      string `json:"name"`
						Part      string `json:"part"`
						Sn        string `json:"sn"`
						ImagePath string `json:"imagePath"`
					} `json:"test_point"`
				} `json:"clearance"`
			} `json:"sub_task"`
		} `json:"main_task"`
	} `json:"data"`
}

func GetTaskContent(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
		}
	}()

	file, _ := os.ReadFile("./data.json")

	var data GetTaskResponse
	json.Unmarshal(file, &data)
	logger.Infof("%+v", data)
	ctx.JSON(http.StatusOK, data)
}
