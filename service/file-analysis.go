package service

import (
	"encoding/binary"
	"fmt"
	"gpk/logger"
	"os"
)

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

		saveBytesToFile(picData, fmt.Sprintf("pic_%d.jpg", i))
		saveBytesToFile(ifrData, fmt.Sprintf("ifr_%d.jpg", i))
	}
}

func saveBytesToFile(bytes []byte, filePath string) {
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
