package api

import (
	"archive/zip"
	"encoding/json"
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
					Id        string             `json:"id"`
					Name      string             `json:"name"`
					Sn        string             `json:"sn"`
					TestPoint []*model.TestPoint `json:"test_point"`
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

	var items []model.AnalysisItem
	for _, subTask := range request.Data.MainTask.SubTask {
		for _, clearance := range subTask.Clearance {
			for _, point := range clearance.TestPoint {
				if point == nil {
					continue
				}
				if p := dbdata.PatrolPointMap.GetPatrolPoint(point.Id); p != nil {
					item := model.AnalysisItem{
						ObjectID:   point.Id,
						ObjectName: point.Name,
						Point:      p,
						LinkPoint:  point,
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
							dat.MakeFile(dirPath, point.Name)
							// err := WriteStringToFile(item.RealFrame, dirPath+"/"+point.Name+"_base64.txt")
							// if err != nil {
							// 	logger.Error(err)
							// }
						}
					}
					items = append(items, item)
				}
			}
		}
	}

	worker := pool.NewAnalysisWorker()
	pool.PAnalysisRunner.Workers.Store(worker.RequestID, worker)
	defer pool.PAnalysisRunner.Workers.Delete(worker.RequestID)
	worker.Work(items)

	// 等待请求结果
	exit := false
	for {
		select {
		// 分析过程超时
		case <-time.After(time.Second * time.Duration(pool.AnalyzeTimeout)):
			logger.Errorf("请求识别唯一标识为%s的图像分析过程超时", worker.RequestID)
			exit = true
		// 处理分析结果
		case result := <-worker.Wc:
			logger.Infof("正在处理分析结果:%s", worker.RequestID)
			logger.Infof("%+v", result)
			for _, item := range items {
				for _, object := range result.ResultsList {
					// 调试模拟结果
					if item.LinkPoint.Id == "0fcd4f0ebf2a432c83c1ee8cff4ec594" {
						item.LinkPoint.Result = 1
						item.LinkPoint.Detail = "模型匹配失败"
						continue
					}

					if item.LinkPoint.Id == "c50a7ee0125a11efa7bc0242ac140065" {
						item.LinkPoint.Result = -1
						item.LinkPoint.Detail = "渗漏油缺陷"
						continue
					}

					if item.LinkPoint.Id == object.ObjectID {
						for _, r := range object.Results {
							logger.Infof("点位分析结果 %s", r.Value)
							switch r.Value {
							case "0":
								item.LinkPoint.Result = 1
								item.LinkPoint.Detail = "模型匹配失败"
							case "-1":
								item.LinkPoint.Result = -1
								item.LinkPoint.Detail = "渗漏油缺陷"
							}
							if r.ResImageBase64 != "" {
								pic_data, err := service.CovertBase64ToPic(r.ResImageBase64)
								if err != nil {
									logger.Errorf("base64图片解析错误:%s", err.Error())
								} else {
									logger.Info("保存识别图片")
									service.SaveBytesToFile(pic_data, "./"+item.Point.Name+"_det.jpg")
								}
							}
						}
						logger.Infof("点位处理结果 %+v", item.LinkPoint)
					}

				}
			}
			exit = true
		}
		if exit {
			logger.Infof("分析结果处理完成 %+v", request)
			break
		}
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
