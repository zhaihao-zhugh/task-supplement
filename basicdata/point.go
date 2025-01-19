package basicdata

import (
	"encoding/json"
	"gpk/logger"
	"supplementary-inspection/model"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

var PatrolPointMap = PointProvider{
	PatrolPointMap: make(map[string]*model.PatrolPoint),
}

type PointProvider struct {
	mu             sync.RWMutex
	PatrolPointMap map[string]*model.PatrolPoint
}

func (p *PointProvider) GetPatrolPoint(id string) *model.PatrolPoint {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.PatrolPointMap[id]
}

func (p *PointProvider) SetPatrolPoint(id string, patrolPoint *model.PatrolPoint) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.PatrolPointMap[id] = patrolPoint
}

func (p *PointProvider) GetPatrolPointMap() map[string]*model.PatrolPoint {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.PatrolPointMap
}

func (p *PointProvider) GetData() {
	// file, err := os.ReadFile("./point.json")
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// for _, v := range gjson.GetBytes(file, "data").Array() {
	// 	var point model.PatrolPoint
	// 	json.Unmarshal([]byte(v.String()), &point)
	// 	p.SetPatrolPoint(point.Guid, &point)
	// }
	url := baseUrl + "patrolpoint?limit=9999"
	for {
		if res, err := GetData(url, nil); err != nil {
			logger.Error(err)
			time.Sleep(10 * time.Second)
		} else {
			if gjson.GetBytes(res, "success").Bool() {
				for _, v := range gjson.GetBytes(res, "data.list").Array() {
					var point model.PatrolPoint
					if err := json.Unmarshal([]byte(v.Raw), &point); err != nil {
						logger.Error(err)
						continue
					}
					p.SetPatrolPoint(point.Guid, &point)
				}
				break
			} else {
				logger.Error(gjson.GetBytes(res, "message").String())
				time.Sleep(10 * time.Second)
			}
		}
	}
}
