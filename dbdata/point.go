package dbdata

import (
	"encoding/json"
	"fmt"
	"gpk/logger"
	"os"
	"supplementary-inspection/model"
	"sync"

	"github.com/tidwall/gjson"
)

var PatrolPointMap = PointProvider{
	PatrolPointMap: make(map[string]*model.PatrolPoint),
}

func init() {
	PatrolPointMap.GetData()
	fmt.Printf("PatrolPointMap: %+v\n", PatrolPointMap.PatrolPointMap)
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
	file, err := os.ReadFile("./point.json")
	if err != nil {
		logger.Error(err)
		return
	}
	for _, v := range gjson.GetBytes(file, "data").Array() {
		var point model.PatrolPoint
		json.Unmarshal([]byte(v.String()), &point)
		p.SetPatrolPoint(point.Guid, &point)
	}
}
