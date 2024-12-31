package model

type Host struct {
	IP   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

type PatrolPoint struct {
	ID              uint32  `json:"id"`
	Guid            string  `json:"guid"`             // 巡视点位ID
	Name            string  `json:"name"`             // 巡视点位名称
	ComponentID     string  `json:"component_id"`     // 设备组件ID
	ComponentName   string  `json:"component_name"`   // 设备组件名称
	DeviceID        string  `json:"device_id"`        // 主设备ID
	DeviceName      string  `json:"device_name"`      // 主设备名称
	DeviceType      uint32  `json:"device_type"`      // 设备类型
	BayID           string  `json:"bay_id"`           // 间隔ID
	BayName         string  `json:"bay_name"`         // 间隔名称
	AreaID          string  `json:"area_id"`          // 区域ID
	AreaName        string  `json:"area_name"`        // 区域名称
	StationID       string  `json:"station_id"`       // 变电站ID
	StationName     string  `json:"station_name"`     // 变电站名称
	StationCode     string  `json:"station_code"`     // 变电站编码
	AppearanceType  uint32  `json:"appearance_type"`  // 外观类型
	HeatingType     uint32  `json:"heating_type"`     // 制热类型
	MeterType       uint32  `json:"meter_type"`       // 仪表类型
	PhaseType       uint32  `json:"phase_type"`       // 相位类型
	RecognitionType uint32  `json:"recognition_type"` // 识别类型
	Unit            string  `json:"unit"`             // 单位
	Config          string  `json:"config"`           // 识别配置
	FileType        uint32  `json:"file_type"`        // 文件类型
	FilePath        string  `json:"file_path"`        // 文件路径
	MaterialID      string  `json:"material_id"`      // 实物ID
	AnalysisType    uint32  `json:"analysis_type"`    // 分析类型
	AnalysisList    string  `json:"analysis_list"`    // 分析列表
	Alarm           bool    `json:"alarm"`            // 是否告警
	AlarmID         string  `json:"alarm_id"`         // 设备告警ID
	LowerValue      float64 `json:"lower_value"`      // 正常范围下限
	UpperValue      float64 `json:"upper_value"`      // 正常范围上限
	AlarmLevel      uint32  `json:"alarm_level"`      // 告警等级
	MaxRange        float64 `json:"max_range"`        // 最大量程
	MinValue        float64 `json:"min_value"`        // 最小刻度
	MaxValue        float64 `json:"max_value"`        // 最大刻度
	LabelAttri      string  `json:"label_attri"`      // 标签属性
	LatestValue     string  `json:"latest_value"`     // 最新结果值
}

// 高德红外测试点位
type TestPoint struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Sn       string `json:"sn"`
	FileName string `json:"filename"`
	Result   int    `json:"result"`
	Detail   string `json:"detail"`
}

// AnalysisItem 向分析主机发送的待分析对象-国网规范
type AnalysisItem struct {
	ObjectID           string          `json:"objectId"`           // 分析请求-分析点位标识
	ObjectName         string          `json:"objectName"`         // 扩展-分析点位名称
	TypeList           []string        `json:"typeList"`           // 分析请求-缺陷类识别可以选多种类型，状态类识别只选一种类型
	ImageUrlList       []string        `json:"imageUrlList"`       // 分析请求-多张图片只用于判别
	ImageNormalUrlPath string          `json:"imageNormalUrlPath"` // 分析请求-可选 判别基准图选填或图像识别配置图
	Config             string          `json:"config,omitempty"`   // 分析请求-扩展-图像识别配置
	StationCode        string          `json:"-"`                  // 变电站编码
	EdgeCode           string          `json:"-"`                  // 边缘节点编码
	TaskResult         *OriginalResult `json:"-"`                  // 原始任务执行结果
	WatchResult        *WatchResult    `json:"-"`                  // 静默监视结果
	Point              *PatrolPoint    `json:"-"`                  // 巡视点位
	CallbackFunc       interface{}     `json:"-"`                  // 回调函数
	Bbox               string          `json:"bbox"`
	TemplateFrame      string          `json:"templateFrame"`
	RealFrame          string          `json:"realFrame"`
	LinkPoint          *TestPoint      `json:"-"`
}

// AnalysisRequest 向分析主机发送的分析请求-国网规范
type AnalysisRequest struct {
	RequestHostIP   string         `json:"requestHostIp"`   // 分析结果反馈ip地址
	RequestHostPort string         `json:"requestHostPort"` // 分析结果返回端口
	RequestID       string         `json:"requestId"`       // 请求分析数据唯一标识UUID
	ObjectList      []AnalysisItem `json:"objectList"`      // 分析对象数组
}

// Rectangle 分析主机返回结果矩形框-国网规范
type Rectangle struct {
	Areas []struct {
		X string `json:"x"`
		Y string `json:"y"`
	} `json:"areas"`
}

// ResultObject 分析主机返回结果对象-国网规范
type ResultObject struct {
	Type           string      `json:"type"`           // 分析类型
	Value          string      `json:"value"`          // 值
	Code           string      `json:"code"`           // 2000=正确 2001=图像数据错误 2002=算法分析失败
	ResImageUrl    string      `json:"resImageUrl"`    // 结果反馈图像url路径
	ResImageBase64 string      `json:"resImageBase64"` // base64图片
	Conf           float32     `json:"conf"`           // 分析结果置信度 范围0-1保留4位小数
	Desc           string      `json:"desc"`           // 分析结果描述
	Pos            []Rectangle `json:"pos"`
}

// ResultObjects 分析主机返回结果对象-国网规范
type ResultObjects struct {
	ObjectID string         `json:"objectId"` // 分析点位标识
	Results  []ResultObject `json:"results"`  // 分析结果
}

// AnalysisResult 分析主机返回结果-国网规范
type AnalysisResult struct {
	RequestID   string          `json:"requestId"`   // 请求分析数据唯一标识
	ResultsList []ResultObjects `json:"resultsList"` // 结果集
}

// CommonMessage 系统消息结构体
type CommonMessage struct {
	Action string `json:"action"`
}

type OriginalTask struct {
	TaskPatrolledID   string `json:"task_patrolled_id"`
	TaskName          string `json:"task_name"`
	TaskCode          string `json:"task_code"`
	TaskState         string `json:"task_state"`
	PlanStartTime     string `json:"plan_start_time"`
	StartTime         string `json:"start_time"`
	TaskProgress      string `json:"task_progress"`
	TaskEstimatedTime string `json:"task_estimated_time"`
	Description       string `json:"description"`
	PatrolType        uint32 `json:"patrol_type"` // 巡视类型 100=联动巡视 101=顺控确认(扩展)
	Priority          uint32 `json:"priority"`    // 优先级(扩展)
	Total             uint32 `json:"total"`       // 点位总数(扩展)
	Executor          string `json:"executor"`    // 任务执行人(扩展)
}

type OriginalTaskMessage struct {
	Action      string         `json:"action"`
	StationCode string         `json:"station_code"`
	EdgeCode    string         `json:"edge_code"`
	Data        []OriginalTask `json:"data"`
}

type FinalTask struct {
	TaskPatrolledID   string `json:"task_patrolled_id"`
	TaskName          string `json:"task_name"`
	TaskCode          string `json:"task_code"`
	TaskState         string `json:"task_state"`
	PlanStartTime     string `json:"plan_start_time"`
	StartTime         string `json:"start_time"`
	TaskProgress      string `json:"task_progress"`
	TaskEstimatedTime string `json:"task_estimated_time"`
	Priority          uint32 `json:"priority"` // 优先级(扩展)
	Description       string `json:"description"`
	StationID         string `json:"station_id"`
	StationName       string `json:"station_name"`
	StationCode       string `json:"station_code"`
	AccessType        uint32 `json:"access_type"` // 接入类型
	EdgeCode          string `json:"edge_code"`   // 边缘节点编码
	TaskType          uint32 `json:"task_type"`   // 0:立即执行  1:定期执行  2:周期执行  3:间隔执行
	PatrolType        uint32 `json:"patrol_type"` // 0:全面巡视 1:例行巡视，2:特殊巡视，3:专项巡视，4:自定义巡视，5:熄灯巡视 6:单点识别 100:联动巡视  101:一键顺控确认
	TaskSource        uint32 `json:"task_source"` // 1=摄像头;2=机器人;3=视频+摄像头
	Total             uint32 `json:"total"`       // 点位总数(扩展)
	Executor          string `json:"executor"`    // 任务执行人(扩展)
	TaskEndTime       string `json:"task_end_time,omitempty"`
}

type FinalTaskMessage struct {
	Action string      `json:"action"`
	Data   []FinalTask `json:"data"`
}

type VoiceResult struct {
	Type      string  `json:"type"`
	Value     string  `json:"value"`
	StartTime float64 `json:"startTime"`
	EndTime   float64 `json:"endTime"`
	Conf      float64 `json:"conf"`
	Desc      string  `json:"desc"`
}

type OriginalResult struct {
	PatrolDeviceName string         `json:"patroldevice_name"`         // 巡视设备名称
	PatrolDeviceCode string         `json:"patroldevice_code"`         // 巡视设备编码
	TaskName         string         `json:"task_name"`                 // 任务名称
	TaskCode         string         `json:"task_code"`                 // 任务编码
	DeviceName       string         `json:"device_name"`               // 设备点位名称
	DeviceID         string         `json:"device_id"`                 // 设备点位ID
	ValueType        string         `json:"value_type"`                // 值类型
	Value            string         `json:"value"`                     // 值
	ValueUnit        string         `json:"value_unit"`                // 值带单位
	Unit             string         `json:"unit"`                      // 单位
	Time             string         `json:"time"`                      // 时间
	RecognitionType  string         `json:"recognition_type"`          // 识别类型
	FileType         string         `json:"file_type"`                 // 采集文件类型 存在多个文件,文件类型用逗号分隔
	FilePath         string         `json:"file_path"`                 // 文件路径 存在多个文件,路径用逗号分隔,并与文件类型一一对应
	Rectangle        string         `json:"rectangle"`                 // 图像框 格式:x1,y1;x2,y2;x3,y3;x4,y4等为图片文件的像素点
	TaskPatrolledID  string         `json:"task_patrolled_id"`         // 巡视任务执行ID
	DataType         string         `json:"data_type"`                 // 巡视结果数据来源 1=摄像机 2=机器人 3=无人机 4=声纹 5=在线监测
	Valid            string         `json:"valid"`                     // 识别结论 0=采集失败 1=成功 2=分析失败
	MonitorId        string         `json:"monitor_id,omitempty"`      // 监控索引号(扩展)
	ShiftFilePath    string         `json:"shift_file_path,omitempty"` // 顺控确认基准图(扩展)
	VoiceResults     *[]VoiceResult `json:"voice_results,omitempty"`   // 声纹分析结果(扩展)
}

type OriginalResultMessage struct {
	Action      string           `json:"action"`
	StationCode string           `json:"station_code"`
	EdgeCode    string           `json:"edge_code"`
	Data        []OriginalResult `json:"data"`
}

type AnalyzedResult struct {
	StationID        string         `json:"station_id"`
	StationName      string         `json:"station_name"`
	StationCode      string         `json:"station_code"`
	EdgeCode         string         `json:"edge_code"`
	PatrolDeviceName string         `json:"patroldevice_name"`       // 巡视设备名称
	PatrolDeviceCode string         `json:"patroldevice_code"`       // 巡视设备编码
	TaskName         string         `json:"task_name"`               // 任务名称
	TaskCode         string         `json:"task_code"`               // 任务编码
	PatrolPointNo    uint32         `json:"patrolpoint_no"`          // 设备点位序号
	PatrolPointName  string         `json:"patrolpoint_name"`        // 设备点位名称
	PatrolPointID    string         `json:"patrolpoint_id"`          // 设备点位ID
	ValueType        string         `json:"value_type"`              // 值类型
	Value            string         `json:"value"`                   // 值(多个用逗号分隔)
	ValueUnit        string         `json:"value_unit"`              // 值带单位(多个用逗号分隔)
	Unit             string         `json:"unit"`                    // 单位
	Time             string         `json:"time"`                    // 时间
	RecognitionType  string         `json:"recognition_type"`        // 识别类型
	FileType         string         `json:"file_type"`               // 文件类型
	FilePath         string         `json:"file_path"`               // 原图路径
	Rectangle        string         `json:"rectangle"`               // 图像框 格式:x1,y1;x2,y2;x3,y3;x4,y4等为图片文件的像素点(多个用逗号分隔)
	TaskPatrolledID  string         `json:"task_patrolled_id"`       // 巡视任务执行ID
	DataType         string         `json:"data_type"`               // 巡视结果数据来源 1=摄像机 2=机器人 3=无人机 4=声纹 5=在线监测
	Valid            string         `json:"valid"`                   // 识别结论 0=采集失败 1=成功 2=分析失败
	Abnormal         bool           `json:"abnormal"`                // 是否异常(缺陷)
	Type             string         `json:"type"`                    // 分析类型(多个用逗号分隔)
	Conf             string         `json:"conf"`                    // 分析结果置信度(多个用逗号分隔)
	PicAnalyzed      string         `json:"pic_analyzed"`            // 分析结果反馈图像
	PicDifferent     string         `json:"pic_different"`           // 判别告警图路径
	PicDiffBase      string         `json:"pic_diff_base"`           // 判别基准图路径
	ComponentID      string         `json:"component_id"`            // 设备组件ID
	ComponentName    string         `json:"component_name"`          // 设备组件名称
	DeviceID         string         `json:"device_id"`               // 主设备ID
	DeviceName       string         `json:"device_name"`             // 主设备名称
	DeviceType       uint32         `json:"device_type"`             // 设备类型
	BayID            string         `json:"bay_id"`                  // 间隔ID
	BayName          string         `json:"bay_name"`                // 间隔名称
	AreaID           string         `json:"area_id"`                 // 区域ID
	AreaName         string         `json:"area_name"`               // 区域名称
	AppearanceType   uint32         `json:"appearance_type"`         // 外观类型
	HeatingType      uint32         `json:"heating_type"`            // 制热类型
	MeterType        uint32         `json:"meter_type"`              // 仪表类型
	PhaseType        uint32         `json:"phase_type"`              // 相位类型
	MaterialID       string         `json:"material_id"`             // 实物ID
	DefectType       string         `json:"defect_type"`             // 缺陷类型 详见 H.6.3.3 缺陷标签(type) 存在多个类型用逗号分隔
	Status           int            `json:"status"`                  // 数据结果(内部) 0:失败 1:成功 2:判别异常 3:告警(缺陷) 4:漏检 10:误报
	Alarm            bool           `json:"alarm"`                   // 是否告警
	AlarmName        string         `json:"alarm_name"`              // 告警名称
	JudgmentType     uint32         `json:"judgment_type"`           // 判断类型
	Expression       string         `json:"expression"`              // 表达式
	LowerValue       float64        `json:"lower_value"`             // 正常范围下限
	UpperValue       float64        `json:"upper_value"`             // 正常范围上限
	AlarmLevel       uint32         `json:"alarm_level"`             // 告警等级
	Remind           bool           `json:"remind"`                  // 是否提醒
	MaxRange         float64        `json:"max_range"`               // 最大量程
	MinValue         float64        `json:"min_value"`               // 最小刻度
	MaxValue         float64        `json:"max_value"`               // 最大刻度
	LabelAttri       string         `json:"label_attri"`             // 标签属性
	VoiceResults     *[]VoiceResult `json:"voice_results,omitempty"` // 声纹分析结果(扩展)
}

type AnalyzedResultMessage struct {
	Action string           `json:"action"`
	Data   []AnalyzedResult `json:"data"`
}

type Alarm struct {
	Name  string `json:"name"`
	Type  int    `json:"type"`  // 6=外观异常
	Level int    `json:"level"` // 1:预警 2:一般 3:严重 4:危急
}

type FinalResult struct {
	StationID        string  `json:"station_id"`
	StationName      string  `json:"station_name"`
	StationCode      string  `json:"station_code"`
	EdgeCode         string  `json:"edge_code"`
	PatrolDeviceName string  `json:"patroldevice_name"`     // 巡视设备名称
	PatrolDeviceCode string  `json:"patroldevice_code"`     // 巡视设备编码
	TaskName         string  `json:"task_name"`             // 任务名称
	TaskCode         string  `json:"task_code"`             // 任务编码
	PatrolPointNo    uint32  `json:"patrolpoint_no"`        // 设备点位序号
	PatrolPointName  string  `json:"patrolpoint_name"`      // 设备点位名称
	PatrolPointID    string  `json:"patrolpoint_id"`        // 设备点位ID
	ValueType        string  `json:"value_type"`            // 值类型
	Value            string  `json:"value"`                 // 值(多个用逗号分隔)
	ValueUnit        string  `json:"value_unit"`            // 值带单位(多个用逗号分隔)
	Unit             string  `json:"unit"`                  // 单位
	Time             string  `json:"time"`                  // 时间
	RecognitionType  string  `json:"recognition_type"`      // 识别类型
	FileType         string  `json:"file_type"`             // 文件类型
	FilePath         string  `json:"file_path"`             // 原图路径
	Rectangle        string  `json:"rectangle"`             // 图像框 格式:x1,y1;x2,y2;x3,y3;x4,y4等为图片文件的像素点(多个用分号分隔)
	TaskPatrolledID  string  `json:"task_patrolled_id"`     // 巡视任务执行ID
	DataType         string  `json:"data_type"`             // 巡视结果数据来源 1=摄像机 2=机器人 3=无人机 4=声纹 5=在线监测
	Valid            string  `json:"valid"`                 // 识别结论 0=采集失败 1=成功 2=分析失败
	Type             string  `json:"type"`                  // 分析类型(多个用逗号分隔)
	Conf             string  `json:"conf"`                  // 分析结果置信度(多个用逗号分隔)
	PicAnalyzed      string  `json:"pic_analyzed"`          // 分析结果反馈图像
	PicDifferent     string  `json:"pic_different"`         // 判别告警图路径
	PicDiffBase      string  `json:"pic_diff_base"`         // 判别基准图路径
	ComponentID      string  `json:"component_id"`          // 设备组件ID
	ComponentName    string  `json:"component_name"`        // 设备组件名称
	DeviceID         string  `json:"device_id"`             // 主设备ID
	DeviceName       string  `json:"device_name"`           // 主设备名称
	DeviceType       uint32  `json:"device_type"`           // 设备类型
	BayID            string  `json:"bay_id"`                // 间隔ID
	BayName          string  `json:"bay_name"`              // 间隔名称
	AreaID           string  `json:"area_id"`               // 区域ID
	AreaName         string  `json:"area_name"`             // 区域名称
	AppearanceType   uint32  `json:"appearance_type"`       // 外观类型
	HeatingType      uint32  `json:"heating_type"`          // 制热类型
	MeterType        uint32  `json:"meter_type"`            // 仪表类型
	PhaseType        uint32  `json:"phase_type"`            // 相位类型
	MaterialID       string  `json:"material_id"`           // 实物ID
	DefectType       string  `json:"defect_type"`           // 缺陷类型 详见 H.6.3.3 缺陷标签(type) 存在多个类型用逗号分隔
	Status           int     `json:"status"`                // 数据结果(内部) 0:失败 1:成功 2:判别异常 3:告警(缺陷) 4:漏检 10:误报
	Alarm            []Alarm `json:"alarm,omitempty"`       // 告警
	LowerValue       float64 `json:"lower_value,omitempty"` // 正常范围下限
	UpperValue       float64 `json:"upper_value,omitempty"` // 正常范围上限
	LabelAttri       string  `json:"label_attri"`           // 标签属性
}

type FinalResultMessage struct {
	Action string        `json:"action"`
	Data   []FinalResult `json:"data"`
}

type WatchResult struct {
	PatrolDeviceCode string `json:"patroldevice_code"` // 巡视设备编码
	DeviceName       string `json:"device_name"`       // 设备点位名称
	DeviceID         string `json:"device_id"`         // 设备点位ID
	Time             string `json:"time"`              // 时间
	Rectangle        string `json:"rectangle"`         // 图像框 格式:x1,y1;x2,y2;x3,y3;x4,y4等为图片文件的像素点 非必填
	FileType         string `json:"file_type"`         // 采集文件类型 2=可见光照片 4=视频
	FilePath         string `json:"file_path"`         // 文件路径
	MonitorType      string `json:"monitor_type"`      // 静默监视类型
}

type WatchResultMessage struct {
	Action string        `json:"action"`
	Data   []WatchResult `json:"data"`
}

type WatchResultAnalyzed struct {
	PatrolDeviceCode string `json:"patroldevice_code"`          // 静默监视设备编码
	PatrolDeviceName string `json:"patroldevice_name"`          // 静默监视设备名称
	AlarmLevel       string `json:"alarm_level"`                // 告警等级 1=预警 2=一般 3=严重 4=危急
	MonitorType      string `json:"monitor_type"`               // 静默监视类型
	FileType         string `json:"file_type"`                  // 告警文件类型 存在多个文件 路径 用逗号分隔 并与文件类型一一对应
	FilePath         string `json:"file_path"`                  // 文件路径 存在多个文件 路径用逗号分隔 并与文件类型一一对应
	Time             string `json:"time"`                       // 时间
	Content          string `json:"content"`                    // 告警描述
	StationID        string `json:"station_id,omitempty"`       // 变电站ID(扩展)
	StationName      string `json:"station_name,omitempty"`     // 变电站名称(扩展)
	PatrolPointID    string `json:"patrolpoint_id,omitempty"`   // 变电站ID(扩展)
	PatrolPointName  string `json:"patrolpoint_name,omitempty"` // 变电站名称(扩展)
	Defect           bool   `json:"defect"`                     // 是否有缺陷
}

type WatchResultAnalyzedMessage struct {
	Action string                `json:"action"`
	Data   []WatchResultAnalyzed `json:"data"`
}

type AdjustRequest struct {
	Data []AdjustObject `json:"data"`
}

type AdjustObject struct {
	CameraID        string `json:"camera_id"`
	PatrolPointCode string `json:"patrolpoint_id"`
	OffsetX         int    `json:"offset_x"`
	OffsetY         int    `json:"offset_y"`
}

// VideoCfmResult 一键顺控视频确认发送到Linkage的内容
type VideoCfmResult struct {
	MonitorId       string `json:"monitor_id"`
	TaskPatrolledID string `json:"task_patrolled_id"`
	PatrolPointID   string `json:"patrolpoint_id"`
	PatrolPointName string `json:"patrolpoint_name"`
	Value           string `json:"value"`
	ValueUnit       string `json:"value_unit"`
	Time            string `json:"time"`
}
