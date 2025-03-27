package basicdata

var BaseUrl string

func Init() {
	PatrolPointMap.GetData()
	TaskMap.GetData()
}
