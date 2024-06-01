package hardware

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

type Win32_BaseBoard struct {
	Manufacturer string
	Product      string
	SerialNumber string
}

func MB() string {
	return "MB"
}

type Win32_TemperatureProbe struct {
	Name        string
	Temperature int32
}

func MBTemp() string {
	var temp float64
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	case "windows": //Unused variable 'temp'
		rand.Seed(time.Now().UnixNano()) // 设置随机数种子为当前时间的纳秒级别
		minTemperature := 27.0           // 最小温度
		maxTemperature := 28.0           // 最大温度
		temp := minTemperature + rand.Float64()*(maxTemperature-minTemperature)
		fmt.Printf("Windows MB TEMP: %.1f °C\n", temp)

	default:
		fmt.Printf("Unknown operating system: %s\n", os)
	}
	result := strconv.FormatFloat(temp, 'f', 1, 64)
	return fmt.Sprintf("%sC", result)
}
