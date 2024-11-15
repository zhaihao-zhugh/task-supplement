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

// type PatrolPoint struct {
// 	ID              uint32  `json:"id"`
// 	Guid            string  `json:"guid"`             // 巡视点位ID
// 	Name            string  `json:"name"`             // 巡视点位名称
// 	ComponentID     string  `json:"component_id"`     // 设备组件ID
// 	ComponentName   string  `json:"component_name"`   // 设备组件名称
// 	DeviceID        string  `json:"device_id"`        // 主设备ID
// 	DeviceName      string  `json:"device_name"`      // 主设备名称
// 	DeviceType      uint32  `json:"device_type"`      // 设备类型
// 	BayID           string  `json:"bay_id"`           // 间隔ID
// 	BayName         string  `json:"bay_name"`         // 间隔名称
// 	AreaID          string  `json:"area_id"`          // 区域ID
// 	AreaName        string  `json:"area_name"`        // 区域名称
// 	StationID       string  `json:"station_id"`       // 变电站ID
// 	StationName     string  `json:"station_name"`     // 变电站名称
// 	StationCode     string  `json:"station_code"`     // 变电站编码
// 	AppearanceType  uint32  `json:"appearance_type"`  // 外观类型
// 	HeatingType     uint32  `json:"heating_type"`     // 制热类型
// 	MeterType       uint32  `json:"meter_type"`       // 仪表类型
// 	PhaseType       uint32  `json:"phase_type"`       // 相位类型
// 	RecognitionType uint32  `json:"recognition_type"` // 识别类型
// 	Unit            string  `json:"unit"`             // 单位
// 	Config          string  `json:"config"`           // 识别配置
// 	FileType        uint32  `json:"file_type"`        // 文件类型
// 	FilePath        string  `json:"file_path"`        // 文件路径
// 	MaterialID      string  `json:"material_id"`      // 实物ID
// 	AnalysisType    uint32  `json:"analysis_type"`    // 分析类型
// 	AnalysisList    string  `json:"analysis_list"`    // 分析列表
// 	Alarm           bool    `json:"alarm"`            // 是否告警
// 	AlarmID         string  `json:"alarm_id"`         // 设备告警ID
// 	LowerValue      float64 `json:"lower_value"`      // 正常范围下限
// 	UpperValue      float64 `json:"upper_value"`      // 正常范围上限
// 	AlarmLevel      uint32  `json:"alarm_level"`      // 告警等级
// 	MaxRange        float64 `json:"max_range"`        // 最大量程
// 	MinValue        float64 `json:"min_value"`        // 最小刻度
// 	MaxValue        float64 `json:"max_value"`        // 最大刻度
// 	LabelAttri      string  `json:"label_attri"`      // 标签属性
// 	LatestValue     string  `json:"latest_value"`     // 最新结果值
// }

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
