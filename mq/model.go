package mq

type TASKITEM struct {
	PatrolpointNo    int         `json:"patrolpoint_no" name:"序号" type:"int"`
	StationID        string      `json:"station_id" name:"变电站id" type:"string"`
	StationName      string      `json:"station_name" name:"变电站名称" type:"string"`
	StationCode      string      `json:"station_code" name:"变电站编码" type:"string"`
	EdgeCode         string      `json:"edge_code" name:"边缘节点编码" type:"string"`
	PatrolDeviceName string      `json:"patroldevice_name" name:"巡视设备名称" type:"string"`
	PatrolDeviceCode string      `json:"patroldevice_code" name:"巡视设备编码" type:"string"`
	TaskName         string      `json:"task_name" name:"任务名称" type:"string"`
	TaskCode         string      `json:"task_code" name:"任务编码" type:"string"`
	PatrolPointName  string      `json:"patrolpoint_name" name:"设备点位名称" type:"string"`
	PatrolPointID    string      `json:"patrolpoint_id" name:"设备点位ID" type:"string"`
	ValueType        string      `json:"value_type" name:"值类型" type:"string"`
	Value            string      `json:"value" name:"巡检值" type:"string"`
	ValueUnit        string      `json:"value_unit" name:"带单位巡检值" type:"string"`
	Unit             string      `json:"unit" name:"单位" type:"string"`
	Time             string      `json:"time" name:"巡视时间" type:"datetime"`
	RecognitionType  string      `json:"recognition_type" name:"识别类型" type:"int" choices:"1=表计读取;2=位置状态识别;3=设备外观查看;4=红外测温;5=声音检测;6=闪烁检测;11=局放超声波检测;12=局放地电压检测;13=局放特高频检测;101=环境温度检测;102=环境湿度检测;103=氧气浓度检测;104=SF6浓度检测"`
	FileType         string      `json:"file_type" name:"文件类型" type:"int" choices:"1=红外图谱;2=可见光照片;3=音频;4=视频"`
	FilePath         string      `json:"file_path" name:"原图路径" type:"image"`
	Rectangle        string      `json:"rectangle" name:"标记区域" type:"string"`
	TaskPatrolledID  string      `json:"task_patrolled_id" name:"任务执行ID" type:"string"`
	DataType         string      `json:"data_type" name:"巡视结果数据来源" type:"int" choices:"1=摄像头;2=机器人;4=无人机;8=声纹;16=在线监测"`
	Valid            string      `json:"valid" name:"巡视结论" type:"int" choices:"0=失败;1=成功;2=判别异常"`
	Type             string      `json:"type" name:"分析类型" type:"string"`
	Conf             string      `json:"conf" name:"分析结果置信度" type:"string"`
	PicAnalyzed      string      `json:"pic_analyzed" name:"分析结果反馈图像" type:"image"`
	PicDifferent     string      `json:"pic_different" name:"判别告警图路径" type:"image"`
	PicDiffBase      string      `json:"pic_diff_base" name:"判别基准图路径" type:"image"`
	ComponentID      string      `json:"component_id" name:"设备组件ID" type:"string"`
	ComponentName    string      `json:"component_name" name:"设备组件名称" type:"string"`
	DeviceID         string      `json:"device_id" name:"主设备ID" type:"string"`
	DeviceName       string      `json:"device_name" name:"主设备名称" type:"string"`
	DeviceType       uint32      `json:"device_type" name:"设备类型" type:"int" choices:"1=油浸式变压器(电抗器);2=断路器;3=组合电器;4=隔离开关;5=开关柜;6=电流互感器;7=电压互感器;8=避雷器;9=并联电容器组;10=干式电抗器;11=串联补偿装置;12=母线及绝缘子;13=穿墙套管;14=消弧线圈;15=高频阻波器;16=耦合电容器;17=高压熔断器;18=中性点隔直(限直)装置;19=接地装置;20=端子箱及检修电源箱;21=站用变压器;22=站用交流电源系统;23=站用直流电源系统;24=设备构架;25=辅助设施;26=土建设施;27=独立避雷针;28=避雷器动作次数表;29=二次屏柜;30=消防系统"`
	BayID            string      `json:"bay_id" name:"间隔ID" type:"string"`
	BayName          string      `json:"bay_name" name:"间隔名称" type:"string"`
	AreaID           string      `json:"area_id" name:"区域ID" type:"string"`
	AreaName         string      `json:"area_name" name:"区域名称" type:"string"`
	AppearanceType   uint32      `json:"appearance_type" name:"辅助设施类型" type:"int" choices:"1=电子围栏;2=红外对射;3=泡沫喷淋;4=消防水泵;5=消防栓;6=消防室;7=设备室;8=照明灯;9=摄像头;10=水位线;11=排水泵;12=沉降监测点"`
	HeatingType      uint32      `json:"heating_type" name:"制热类型" type:"int" choices:"0=电流致热型;1=电压致热型;2=综合致热型"`
	MeterType        uint32      `json:"meter_type" name:"仪表类型" type:"int" choices:"1=油位表;2=避雷器动作次数表;3=泄漏电流表;4=SF6压力表;5=液压表;6=开关动作次数表;7=油温表;8=档位表;9=气压表"`
	PhaseType        uint32      `json:"phase_type" name:"相位类型" type:"int" choices:"0=三相;1=A相;2=B相;3=C相"`
	MaterialID       string      `json:"material_id" name:"实物ID" type:"string"`
	Status           int         `json:"status" name:"数据状态" type:"int" choices:"0=失败;1=正常;2=判别异常;3=告警(缺陷);4=漏检;10=误报"`
	Alarm            interface{} `json:"alarm,omitempty"`
	Threshold        interface{} `json:"threshold,omitempty" name:"阈值" type:"string"`
	CheckMessage     string      `json:"check_message" name:"审核意见" type:"string"`
	CheckStatus      int         `json:"check_status" name:"审核状态" type:"int" choices:"0=未审核;1=已审核"`
	Checkman         string      `json:"check_man" name:"审核人" type:"string"`
	CheckTime        string      `json:"check_time,omitempty" name:"审核时间" type:"datetime"`
	CorrectedValue   string      `json:"corrected_value,omitempty" name:"修正值" type:"string"`
	CorrectedStatus  int         `json:"corrected_status" name:"修正状态" type:"int" choices:"-1=未修正;0=失败;1=正常;2=判别异常;3=告警(缺陷);4=漏检;10=误报"`
	Remind           bool        `json:"remind"` // 是否提醒
	Skip             bool        `json:"skip"`
	Detail           string      `json:"detail"`
	DefectType       string      `json:"defect_type"`
	LabelAttri       string      `json:"label_attri"`
	VoiceResults     interface{} `json:"voice_results,omitempty"`
}
