package api

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"gpk/logger"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"supplementary-inspection/dbdata"
	"supplementary-inspection/model"
	"supplementary-inspection/pool"
	"supplementary-inspection/service"
	"time"

	"github.com/gin-gonic/gin"
)

type File interface {
	io.Reader
}

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
	file := form.File["files"][0]
	dirPath := "/store/ftp/tmp/" + strings.Split(file.Filename, ".")[0]
	filePath := dirPath + "/" + file.Filename
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	zr, err := zip.OpenReader(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer zr.Close()
	for _, f := range zr.File {
		logger.Infof("解压到文件:%+v\n", f)
		// 直接读文件的内容
		reader, err := f.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer reader.Close()
		// dat := service.AnalyzeDatFile(reader)
		// dat.MakeFile("/store/ftp/tmp", strings.Split(f.Name, ".")[0])

		// 保存文件
		err = SaveFile(reader, dirPath+"/"+f.Name)
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
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	for _, subTask := range request.Data.MainTask.SubTask {
		for _, clearance := range subTask.Clearance {
			for _, point := range clearance.TestPoint {
				if p := dbdata.PatrolPointMap.GetPatrolPoint(point.Id); p != nil {
					item := model.AnalysisItem{
						ObjectID:   point.Id,
						ObjectName: point.Name,
						Point:      p,
					}
					if p.AnalysisList != "" {
						item.TypeList = strings.Split(p.AnalysisList, ",")
					}
					if p.FilePath != "" {
						if file, err := os.ReadFile(p.FilePath); err == nil {
							item.TemplateFrame = service.CovertPicToBase64(file)
						} else {
							logger.Error(err)
						}
					}

					if point.FileName != "" {
						if dat := service.AnalyzeDatFileByFilepath(dirPath + "/" + point.FileName); dat != nil {
							item.RealFrame = service.CovertPicToBase64(dat.GetPicData(0))
							if item.TemplateFrame == "" {
								item.TemplateFrame = item.RealFrame
							}
							err := WriteStringToFile(item.RealFrame, dirPath+"/"+point.Name+"_base64.txt")
							if err != nil {
								logger.Error(err)
							}
						}
					}
					pool.PAnalysisRunner.Append(item)
				}
			}
		}
	}

	// 等待请求结果
	select {
	// 分析过程超时
	case <-time.After(time.Second * time.Duration(pool.AnalyzeTimeout)):
		// for _, item := range items {
		// 	item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析过程超时")
		// 	logger.Errorf("###分析过程超时:%s", item.Point.Name)
		// }
		logger.Errorf("请求识别唯一标识为%s的图像分析过程超时", pool.worker.RequestID)
		res <- fmt.Errorf("分析主机 图像分析超时")
		return
	// 处理分析结果
	case result := <-worker.Wc:
		logger.Infof("正在处理分析结果:%s", worker.RequestID)
		for _, item := range items {
			// exist := false
			for _, object := range result.ResultsList {

				// if object.ObjectID == item.ObjectID {
				// 	exist = true
				// 	// item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(&object, &item, "")
				// }

				if object.ResImageBase64 != "" {
					pic_data, err := service.CovertBase64ToPic(object.ResImageBase64)
					if err != nil {
						logger.Errorf("base64图片解析错误:%s", err.Error())
					} else {
						logger.Info("保存识别图片")
						service.SaveBytesToFile(pic_data, "./"+item.Point.Name+"_det.jpg")
					}
				}
			}
			// if !exist {
			// 	// item.CallbackFunc.(func(*model.ResultObjects, *model.AnalysisItem, string))(nil, &item, "分析主机漏检")
			// 	logger.Errorf("###分析主机漏检:%s", item.Point.Name)
			// }
		}

		res <- nil
		logger.Info("分析结果处理完成")
		return
	}

	ctx.JSON(http.StatusOK, request)
}

func SaveFile(file File, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}

func WriteStringToFile(str_data string, dst string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(str_data)
	if err != nil {
		return err
	}
	return nil
}
