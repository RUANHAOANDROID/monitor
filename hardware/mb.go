package hardware

import (
	"fmt"
	"github.com/yusufpapurcu/wmi"
	"log"
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
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	case "windows": //Unused variable 'temp'
		var dst []Win32_BaseBoard
		query := "SELECT Manufacturer, Product, SerialNumber,Temp FROM Win32_BaseBoard"
		if err := wmi.Query(query, &dst); err != nil {
			log.Fatal(err)
		}
		for _, item := range dst {
			fmt.Printf("Manufacturer: %s\n", item.Manufacturer)
			fmt.Printf("Product: %s\n", item.Product)
			fmt.Printf("SerialNumber: %s\n", item.SerialNumber)
		}
	default:
		fmt.Printf("Unknown operating system: %s\n", os)
	}
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
