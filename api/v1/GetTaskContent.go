package api

import (
	"gpk/logger"
	"net/http"
	"strconv"
	"strings"
	"supplementary-inspection/basicdata"

	"github.com/gin-gonic/gin"
)

type GetTaskResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		MainTask []Task `json:"main_task"`
	} `json:"data"`
}

// 任务
type Task struct {
	Id      string     `json:"id"`
	Type    string     `json:"type"`
	Name    string     `json:"name"`
	SubTask []*SubTask `json:"sub_task"`
}

// 区域
type SubTask struct {
	Id        string       `json:"id"`
	Type      string       `json:"type"`
	Name      string       `json:"name"`
	Clearance []*Clearance `json:"clearance"`
}

// 间隔
type Clearance struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	Sn        string       `json:"sn"`
	TestPoint []*TestPoint `json:"test_point"`
}

// 点位
type TestPoint struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Part      string `json:"part"`
	Sn        string `json:"sn"`
	ImagePath string `json:"imagePath"`
	MaskUrl   string `json:"maskUrl"`
}

func GetTaskContent(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
		}
	}()

	var res GetTaskResponse
	// file, _ := os.ReadFile("./data.json")
	// json.Unmarshal(file, &data)

	// basicdata.Init()
	taskList := basicdata.TaskMap.GetTaskMap()
	for _, v := range taskList {
		if v.DeviceLevel == 3 && v.DeviceList != "" {
			task := Task{
				Id:   v.GUID,
				Type: "盲点检测",
				Name: v.Name,
			}
			clearanceMap := make(map[string]*Clearance)
			subtaskMap := make(map[string]*SubTask)

			points := strings.Split(v.DeviceList, ",")
			for i, p := range points {
				if patrol_point := basicdata.PatrolPointMap.GetPatrolPoint(p); patrol_point != nil {
					point := TestPoint{
						Id:        patrol_point.Guid,
						Name:      patrol_point.Name,
						Part:      patrol_point.Component.Name,
						Sn:        strconv.Itoa(i + 1),
						ImagePath: patrol_point.ImageUrl,
						MaskUrl:   patrol_point.MaskUrl,
					}

					var c_p *Clearance
					if c, ok := clearanceMap[patrol_point.Bay.Guid]; ok {
						c.TestPoint = append(c.TestPoint, &point)
						c_p = c
					} else {
						c_p = &Clearance{
							Id:        patrol_point.Bay.Guid,
							Name:      patrol_point.Bay.Name,
							Sn:        strconv.Itoa(len(clearanceMap) + 1),
							TestPoint: []*TestPoint{&point},
						}
						clearanceMap[patrol_point.Bay.Guid] = c_p
					}

					if _, ok := subtaskMap[patrol_point.Area.Guid]; !ok {
						subtaskMap[patrol_point.Area.Guid] = &SubTask{
							Id:        patrol_point.Area.Guid,
							Name:      patrol_point.Area.Name,
							Type:      "区域",
							Clearance: []*Clearance{c_p},
						}
					}
				}
			}

			for _, sub := range subtaskMap {
				task.SubTask = append(task.SubTask, sub)
			}
			res.Data.MainTask = append(res.Data.MainTask, task)
		}
	}

	res.Code = 200
	logger.Infof("%+v", res)
	ctx.JSON(http.StatusOK, res)
}
