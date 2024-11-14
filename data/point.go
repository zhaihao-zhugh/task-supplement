package data

import "sync"

var PatrolPointMap = PointProvider{
	PatrolPointMap: make(map[string]*PatrolPoint),
}

type PointProvider struct {
	mu             sync.RWMutex
	PatrolPointMap map[string]*PatrolPoint
}

type PatrolPoint struct {
	ID              string `json:"id"`
	Guid            string `json:"guid"`             // 巡视点位ID
	Name            string `json:"name"`             // 巡视点位名称
	ComponentID     string `json:"component_id"`     // 设备组件ID
	ComponentName   string `json:"component_name"`   // 设备组件名称
	DeviceID        string `json:"device_id"`        // 主设备ID
	DeviceName      string `json:"device_name"`      // 主设备名称
	DeviceType      uint32 `json:"device_type"`      // 设备类型
	BayID           string `json:"bay_id"`           // 间隔ID
	BayName         string `json:"bay_name"`         // 间隔名称
	AreaID          string `json:"area_id"`          // 区域ID
	AreaName        string `json:"area_name"`        // 区域名称
	StationID       string `json:"station_id"`       // 变电站ID
	StationName     string `json:"station_name"`     // 变电站名称
	StationCode     string `json:"station_code"`     // 变电站编码
	AppearanceType  uint32 `json:"appearance_type"`  // 外观类型
	HeatingType     uint32 `json:"heating_type"`     // 制热类型
	MeterType       uint32 `json:"meter_type"`       // 仪表类型
	PhaseType       uint32 `json:"phase_type"`       // 相位类型
	RecognitionType uint32 `json:"recognition_type"` // 识别类型
	FileType        uint32 `json:"file_type"`        // 文件类型
	FilePath        string `json:"file_path"`        // 文件路径
	MaterialID      string `json:"material_id"`      // 实物ID
}

func (p *PointProvider) GetPatrolPoint(id string) *PatrolPoint {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.PatrolPointMap[id]
}

func (p *PointProvider) SetPatrolPoint(id string, patrolPoint *PatrolPoint) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.PatrolPointMap[id] = patrolPoint
}

func (p *PointProvider) GetPatrolPointMap() map[string]*PatrolPoint {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.PatrolPointMap
}

func (p *PointProvider) GetData() map[string]*PatrolPoint {

}
