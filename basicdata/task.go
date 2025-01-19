package basicdata

import (
	"encoding/json"
	"gpk/logger"
	"supplementary-inspection/model"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

var TaskMap = TaskProvider{
	TaskMap: make(map[string]*model.Task),
}

type TaskProvider struct {
	mu      sync.RWMutex
	TaskMap map[string]*model.Task
}

func (t *TaskProvider) GetTask(id string) *model.Task {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.TaskMap[id]
}

func (t *TaskProvider) SetTask(id string, task *model.Task) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.TaskMap[id] = task
}

func (t *TaskProvider) GetTaskMap() map[string]*model.Task {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.TaskMap
}

func (t *TaskProvider) GetData() {
	url := baseUrl + "task?limit=9999"
	for {
		if res, err := GetData(url, nil); err != nil {
			logger.Error(err)
			time.Sleep(10 * time.Second)
		} else {
			if gjson.GetBytes(res, "success").Bool() {
				for _, v := range gjson.GetBytes(res, "data.list").Array() {
					var task model.Task
					if err := json.Unmarshal([]byte(v.Raw), &task); err != nil {
						logger.Error(err)
						continue
					}
					t.SetTask(task.GUID, &task)
				}
				break
			} else {
				logger.Error(gjson.GetBytes(res, "message").String())
				time.Sleep(10 * time.Second)
			}
		}
	}
}
