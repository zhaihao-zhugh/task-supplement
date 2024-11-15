package service

import (
	"encoding/binary"
	"fmt"
	"gpk/logger"
	"os"
)

type File interface {
	Read(p []byte) (n int, err error)
}

func LoadFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	//通用数据格式框架
	var length uint32
	var version [4]byte
	var createTime uint64
	var stationName = make([]byte, 118)
	var stationCode = make([]byte, 42)
	var weather uint8
	var temperature float32
	var humidity uint8
	var deviceManufacturer = make([]byte, 32)
	var deviceModel = make([]byte, 32)
	var deviceVersion [4]byte
	var deviceSerialNum = make([]byte, 32)
	var frequency float32
	var fileCount uint16
	var longitude float64
	var latitude float64
	var altitude int32
	var skip = make([]byte, 204)

	binary.Read(file, binary.LittleEndian, &length)             // 文件长度
	binary.Read(file, binary.LittleEndian, &version)            // 规范版本号
	binary.Read(file, binary.LittleEndian, &createTime)         // 文件生成时间
	binary.Read(file, binary.LittleEndian, &stationName)        // 站点名称
	binary.Read(file, binary.LittleEndian, &stationCode)        // 站点编码
	binary.Read(file, binary.LittleEndian, &weather)            // 天气
	binary.Read(file, binary.LittleEndian, &temperature)        // 温度
	binary.Read(file, binary.LittleEndian, &humidity)           // 湿度
	binary.Read(file, binary.LittleEndian, &deviceManufacturer) // 仪器厂家
	binary.Read(file, binary.LittleEndian, &deviceModel)        // 仪器型号
	binary.Read(file, binary.LittleEndian, &deviceVersion)      // 仪器版本号
	binary.Read(file, binary.LittleEndian, &deviceSerialNum)    // 仪器序列号
	binary.Read(file, binary.LittleEndian, &frequency)          // 系统频率
	binary.Read(file, binary.LittleEndian, &fileCount)          // 图谱数量
	binary.Read(file, binary.LittleEndian, &longitude)          // 经度
	binary.Read(file, binary.LittleEndian, &latitude)           // 纬度
	binary.Read(file, binary.LittleEndian, &altitude)           // 海拔
	binary.Read(file, binary.LittleEndian, &skip)               // 预留

	logger.Infof("文件长度:%v", length)
	logger.Infof("规范版本号:%v.%v.%v.%v", version[0], version[1], version[2], version[3])
	logger.Infof("文件生成时间:%v", createTime)
	logger.Infof("站点名称:%v", string(stationName))
	logger.Infof("站点编码:%v", string(stationCode))
	logger.Infof("天气:%v", weather)
	logger.Infof("温度:%v", temperature)
	logger.Infof("湿度:%v", humidity)
	logger.Infof("仪器厂家:%v", string(deviceManufacturer))
	logger.Infof("仪器型号:%v", string(deviceModel))
	logger.Infof("仪器版本号:%v.%v.%v.%v", deviceVersion[0], deviceVersion[1], deviceVersion[2], deviceVersion[3])
	logger.Infof("仪器序列号:%v", string(deviceSerialNum))
	logger.Infof("系统频率:%v", frequency)
	logger.Infof("图谱数量:%v", fileCount)
	logger.Infof("经度:%v", longitude)
	logger.Infof("纬度:%v", latitude)
	logger.Infof("海拔:%v", altitude)

	for i := 0; i < int(fileCount); i++ {
		// 红外图谱数据格式
		var datType uint8                  // 检测数据类型
		var datLength uint32               // 图谱数据长度
		var datCreateTime uint64           // 图谱生成时间
		var datNature uint8                // 图谱性质
		var targetName = make([]byte, 118) // 被测设备名称
		var targetCode = make([]byte, 42)  // 被测设备编码
		var pointName = make([]byte, 128)  // 测点名称
		var pointCode = make([]byte, 32)   // 测点编码
		var channel int16                  // 检测通道标志
		var saveDataType uint8             // 存储数据类型
		var temperatureUnit uint8          // 温度单位
		var temperatureWidth int32         // 温度点阵宽度
		var temperatureHeight int32        // 温度点阵高度
		var picSize uint32                 // 可见光照片数据长度
		var ifrSize uint32                 // 红外照片数据长度
		var emissivity float32             // 发射率
		var distance float32               // 测试距离
		var envTemperature float32         // 大气温度
		var envHumidity uint8              // 相对湿度
		var refTemperature float32         // 反射温度
		var upperLimit float32             // 温宽上限
		var lowerLimit float32             // 温宽下限
		var datSkip = make([]byte, 133)    // 预留

		binary.Read(file, binary.LittleEndian, &datType)
		binary.Read(file, binary.LittleEndian, &datLength)
		binary.Read(file, binary.LittleEndian, &datCreateTime)
		binary.Read(file, binary.LittleEndian, &datNature)
		binary.Read(file, binary.LittleEndian, &targetName)
		binary.Read(file, binary.LittleEndian, &targetCode)
		binary.Read(file, binary.LittleEndian, &pointName)
		binary.Read(file, binary.LittleEndian, &pointCode)
		binary.Read(file, binary.LittleEndian, &channel)
		binary.Read(file, binary.LittleEndian, &saveDataType)
		binary.Read(file, binary.LittleEndian, &temperatureUnit)
		binary.Read(file, binary.LittleEndian, &temperatureWidth)
		binary.Read(file, binary.LittleEndian, &temperatureHeight)
		binary.Read(file, binary.LittleEndian, &picSize)
		binary.Read(file, binary.LittleEndian, &ifrSize)
		binary.Read(file, binary.LittleEndian, &emissivity)
		binary.Read(file, binary.LittleEndian, &distance)
		binary.Read(file, binary.LittleEndian, &envTemperature)
		binary.Read(file, binary.LittleEndian, &envHumidity)
		binary.Read(file, binary.LittleEndian, &refTemperature)
		binary.Read(file, binary.LittleEndian, &upperLimit)
		binary.Read(file, binary.LittleEndian, &lowerLimit)
		binary.Read(file, binary.LittleEndian, &datSkip)

		logger.Infof("检测数据类型:%v", datType)
		logger.Infof("图谱数据长度:%v", datLength)
		logger.Infof("图谱生成时间:%v", datCreateTime)
		logger.Infof("图谱性质:%v", datNature)
		logger.Infof("被测设备名称:%v", string(targetName))
		logger.Infof("被测设备编码:%v", string(targetCode))
		logger.Infof("测点名称:%v", string(pointName))
		logger.Infof("测点编码:%v", string(pointCode))
		logger.Infof("检测通道标志:%v", channel)
		logger.Infof("存储数据类型:%v", saveDataType)
		logger.Infof("温度单位:%v", temperatureUnit)
		logger.Infof("温度点阵宽度:%v", temperatureWidth)
		logger.Infof("温度点阵高度:%v", temperatureHeight)
		logger.Infof("可见光照片数据长度:%v", picSize)
		logger.Infof("红外照片数据长度:%v", ifrSize)
		logger.Infof("发射率:%v", emissivity)
		logger.Infof("测试距离:%v", distance)
		logger.Infof("大气温度:%v", envTemperature)
		logger.Infof("相对湿度:%v", envHumidity)
		logger.Infof("反射温度:%v", refTemperature)
		logger.Infof("温宽上限:%v", upperLimit)
		logger.Infof("温宽下限:%v", lowerLimit)

		switch saveDataType {
		case 0x02:
			var datData = make([]byte, temperatureWidth*temperatureHeight) // 红外图谱数据
			binary.Read(file, binary.LittleEndian, &datData)
		default:
			var datData = make([]byte, 4*temperatureWidth*temperatureHeight) // 红外图谱数据
			binary.Read(file, binary.LittleEndian, &datData)
		}

		var picData = make([]byte, picSize)
		var ifrData = make([]byte, ifrSize)

		binary.Read(file, binary.LittleEndian, &picData)
		binary.Read(file, binary.LittleEndian, &ifrData)

		SaveBytesToFile(picData, fmt.Sprintf("pic_%d.jpg", i))
		SaveBytesToFile(ifrData, fmt.Sprintf("ifr_%d.jpg", i))
	}
}

func SaveBytesToFile(bytes []byte, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		logger.Errorf("Error creating file: %v", err.Error())
		return
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		logger.Errorf("Error writing to file: %v", err.Error())
		return
	}
}

func AnalyzeDatFileByFilepath(filePath string) *GaoDeDat {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer file.Close()

	dat := new(GaoDeDat)
	binary.Read(file, binary.LittleEndian, &dat.length)             // 文件长度
	binary.Read(file, binary.LittleEndian, &dat.version)            // 规范版本号
	binary.Read(file, binary.LittleEndian, &dat.createTime)         // 文件生成时间
	binary.Read(file, binary.LittleEndian, &dat.stationName)        // 站点名称
	binary.Read(file, binary.LittleEndian, &dat.stationCode)        // 站点编码
	binary.Read(file, binary.LittleEndian, &dat.weather)            // 天气
	binary.Read(file, binary.LittleEndian, &dat.temperature)        // 温度
	binary.Read(file, binary.LittleEndian, &dat.humidity)           // 湿度
	binary.Read(file, binary.LittleEndian, &dat.deviceManufacturer) // 仪器厂家
	binary.Read(file, binary.LittleEndian, &dat.deviceModel)        // 仪器型号
	binary.Read(file, binary.LittleEndian, &dat.deviceVersion)      // 仪器版本号
	binary.Read(file, binary.LittleEndian, &dat.deviceSerialNum)    // 仪器序列号
	binary.Read(file, binary.LittleEndian, &dat.frequency)          // 系统频率
	binary.Read(file, binary.LittleEndian, &dat.fileCount)          // 图谱数量
	binary.Read(file, binary.LittleEndian, &dat.longitude)          // 经度
	binary.Read(file, binary.LittleEndian, &dat.latitude)           // 纬度
	binary.Read(file, binary.LittleEndian, &dat.altitude)           // 海拔
	binary.Read(file, binary.LittleEndian, &dat.skip)               // 预留

	for i := 0; i < int(dat.fileCount); i++ {
		f := new(GaoDeFile)
		binary.Read(file, binary.LittleEndian, &f.datType)
		binary.Read(file, binary.LittleEndian, &f.datLength)
		binary.Read(file, binary.LittleEndian, &f.datCreateTime)
		binary.Read(file, binary.LittleEndian, &f.datNature)
		binary.Read(file, binary.LittleEndian, &f.targetName)
		binary.Read(file, binary.LittleEndian, &f.targetCode)
		binary.Read(file, binary.LittleEndian, &f.pointName)
		binary.Read(file, binary.LittleEndian, &f.pointCode)
		binary.Read(file, binary.LittleEndian, &f.channel)
		binary.Read(file, binary.LittleEndian, &f.saveDataType)
		binary.Read(file, binary.LittleEndian, &f.temperatureUnit)
		binary.Read(file, binary.LittleEndian, &f.temperatureWidth)
		binary.Read(file, binary.LittleEndian, &f.temperatureHeight)
		binary.Read(file, binary.LittleEndian, &f.picSize)
		binary.Read(file, binary.LittleEndian, &f.ifrSize)
		binary.Read(file, binary.LittleEndian, &f.emissivity)
		binary.Read(file, binary.LittleEndian, &f.distance)
		binary.Read(file, binary.LittleEndian, &f.envTemperature)
		binary.Read(file, binary.LittleEndian, &f.envHumidity)
		binary.Read(file, binary.LittleEndian, &f.refTemperature)
		binary.Read(file, binary.LittleEndian, &f.upperLimit)
		binary.Read(file, binary.LittleEndian, &f.lowerLimit)
		binary.Read(file, binary.LittleEndian, &f.datSkip)

		switch f.saveDataType {
		case 0x02:
			var datData = make([]byte, f.temperatureWidth*f.temperatureHeight) // 红外图谱数据
			binary.Read(file, binary.LittleEndian, &datData)
		default:
			var datData = make([]byte, 4*f.temperatureWidth*f.temperatureHeight) // 红外图谱数据
			binary.Read(file, binary.LittleEndian, &datData)
		}

		f.PicData = make([]byte, f.picSize)
		f.IfrData = make([]byte, f.ifrSize)

		binary.Read(file, binary.LittleEndian, &f.PicData)
		binary.Read(file, binary.LittleEndian, &f.IfrData)

		dat.files = append(dat.files, f)
	}

	return dat
}

func AnalyzeDatFile(file File) *GaoDeDat {
	dat := new(GaoDeDat)
	binary.Read(file, binary.LittleEndian, &dat.length)             // 文件长度
	binary.Read(file, binary.LittleEndian, &dat.version)            // 规范版本号
	binary.Read(file, binary.LittleEndian, &dat.createTime)         // 文件生成时间
	binary.Read(file, binary.LittleEndian, &dat.stationName)        // 站点名称
	binary.Read(file, binary.LittleEndian, &dat.stationCode)        // 站点编码
	binary.Read(file, binary.LittleEndian, &dat.weather)            // 天气
	binary.Read(file, binary.LittleEndian, &dat.temperature)        // 温度
	binary.Read(file, binary.LittleEndian, &dat.humidity)           // 湿度
	binary.Read(file, binary.LittleEndian, &dat.deviceManufacturer) // 仪器厂家
	binary.Read(file, binary.LittleEndian, &dat.deviceModel)        // 仪器型号
	binary.Read(file, binary.LittleEndian, &dat.deviceVersion)      // 仪器版本号
	binary.Read(file, binary.LittleEndian, &dat.deviceSerialNum)    // 仪器序列号
	binary.Read(file, binary.LittleEndian, &dat.frequency)          // 系统频率
	binary.Read(file, binary.LittleEndian, &dat.fileCount)          // 图谱数量
	binary.Read(file, binary.LittleEndian, &dat.longitude)          // 经度
	binary.Read(file, binary.LittleEndian, &dat.latitude)           // 纬度
	binary.Read(file, binary.LittleEndian, &dat.altitude)           // 海拔
	binary.Read(file, binary.LittleEndian, &dat.skip)               // 预留

	for i := 0; i < int(dat.fileCount); i++ {
		f := new(GaoDeFile)
		binary.Read(file, binary.LittleEndian, &f.datType)
		binary.Read(file, binary.LittleEndian, &f.datLength)
		binary.Read(file, binary.LittleEndian, &f.datCreateTime)
		binary.Read(file, binary.LittleEndian, &f.datNature)
		binary.Read(file, binary.LittleEndian, &f.targetName)
		binary.Read(file, binary.LittleEndian, &f.targetCode)
		binary.Read(file, binary.LittleEndian, &f.pointName)
		binary.Read(file, binary.LittleEndian, &f.pointCode)
		binary.Read(file, binary.LittleEndian, &f.channel)
		binary.Read(file, binary.LittleEndian, &f.saveDataType)
		binary.Read(file, binary.LittleEndian, &f.temperatureUnit)
		binary.Read(file, binary.LittleEndian, &f.temperatureWidth)
		binary.Read(file, binary.LittleEndian, &f.temperatureHeight)
		binary.Read(file, binary.LittleEndian, &f.picSize)
		binary.Read(file, binary.LittleEndian, &f.ifrSize)
		binary.Read(file, binary.LittleEndian, &f.emissivity)
		binary.Read(file, binary.LittleEndian, &f.distance)
		binary.Read(file, binary.LittleEndian, &f.envTemperature)
		binary.Read(file, binary.LittleEndian, &f.envHumidity)
		binary.Read(file, binary.LittleEndian, &f.refTemperature)
		binary.Read(file, binary.LittleEndian, &f.upperLimit)
		binary.Read(file, binary.LittleEndian, &f.lowerLimit)
		binary.Read(file, binary.LittleEndian, &f.datSkip)

		switch f.saveDataType {
		case 0x02:
			var datData = make([]byte, f.temperatureWidth*f.temperatureHeight) // 红外图谱数据
			binary.Read(file, binary.LittleEndian, &datData)
		default:
			var datData = make([]byte, 4*f.temperatureWidth*f.temperatureHeight) // 红外图谱数据
			binary.Read(file, binary.LittleEndian, &datData)
		}

		f.PicData = make([]byte, f.picSize)
		f.IfrData = make([]byte, f.ifrSize)

		binary.Read(file, binary.LittleEndian, &f.PicData)
		binary.Read(file, binary.LittleEndian, &f.IfrData)

		dat.files = append(dat.files, f)
	}

	return dat
}

type GaoDeDat struct {
	length             uint32
	version            [4]byte
	createTime         uint64
	stationName        [118]byte
	stationCode        [42]byte
	weather            uint8
	temperature        float32
	humidity           uint8
	deviceManufacturer [32]byte
	deviceModel        [32]byte
	deviceVersion      [4]byte
	deviceSerialNum    [32]byte
	frequency          float32
	fileCount          uint16
	longitude          float64
	latitude           float64
	altitude           int32
	skip               [204]byte
	files              []*GaoDeFile
}

func (g *GaoDeDat) MakeFile(filepath, filename string) {
	for _, f := range g.files {
		f.MakePicFile(fmt.Sprintf("%s/%s_pic.jpg", filepath, filename))
		f.MakeIfrFile(fmt.Sprintf("%s/%s_ifr.jpg", filepath, filename))
	}
}

func (g *GaoDeDat) GetPicData(i int) []byte {
	if i+1 > len(g.files) {
		return nil
	}
	return g.files[i].PicData
}

func (g *GaoDeDat) GetIfrData(i int) []byte {
	if i+1 > len(g.files) {
		return nil
	}
	return g.files[i].IfrData
}

type GaoDeFile struct {
	datType           uint8     // 检测数据类型
	datLength         uint32    // 图谱数据长度
	datCreateTime     uint64    // 图谱生成时间
	datNature         uint8     // 图谱性质
	targetName        [118]byte // 被测设备名称
	targetCode        [42]byte  // 被测设备编码
	pointName         [128]byte // 测点名称
	pointCode         [32]byte  // 测点编码
	channel           int16     // 检测通道标志
	saveDataType      uint8     // 存储数据类型
	temperatureUnit   uint8     // 温度单位
	temperatureWidth  int32     // 温度点阵宽度
	temperatureHeight int32     // 温度点阵高度
	picSize           uint32    // 可见光照片数据长度
	ifrSize           uint32    // 红外照片数据长度
	emissivity        float32   // 发射率
	distance          float32   // 测试距离
	envTemperature    float32   // 大气温度
	envHumidity       uint8     // 相对湿度
	refTemperature    float32   // 反射温度
	upperLimit        float32   // 温宽上限
	lowerLimit        float32   // 温宽下限
	datSkip           [133]byte // 预留
	PicData           []byte
	IfrData           []byte
}

func (f *GaoDeFile) MakePicFile(filepath string) {
	SaveBytesToFile(f.PicData, filepath)
}

func (f *GaoDeFile) MakeIfrFile(filepath string) {
	SaveBytesToFile(f.IfrData, filepath)
}
