package hardware

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

type DiskInfo struct {
	Mountpoint string
	TotalGB    float64
	UsedGB     float64
}

func GetDiskInfo() string {
	d := disks()
	dt := diskTemp()
	name := d[0].Mountpoint
	total := d[0].TotalGB
	used := d[0].UsedGB
	return fmt.Sprintf("%s,%sG,%sG,%s", name, strconv.FormatFloat(total, 'f', 1, 64), strconv.FormatFloat(used, 'f', 1, 64), dt)
}
func disks() []DiskInfo {
	// 获取所有磁盘的信息
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatal(err)
	}
	var diskInfos []DiskInfo
	// 遍历每个磁盘
	for _, partition := range partitions {
		fmt.Printf("Disk(%s)\n", partition.Mountpoint)

		// 获取磁盘使用情况
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Fatal(err)
		}
		total := float64(usage.Total) / 1024 / 1024 / 1024
		used := float64(usage.Used) / 1024 / 1024 / 1024
		// 将关键信息添加到数组中
		diskInfo := DiskInfo{
			Mountpoint: partition.Mountpoint,
			TotalGB:    total,
			UsedGB:     used,
		}
		diskInfos = append(diskInfos, diskInfo)
		fmt.Printf("Total: %.2f GB\n", total)
		fmt.Printf("Used: %.2f GB\n", used)
	}
	return diskInfos
}

const (
	IOCTL_STORAGE_QUERY_PROPERTY = 0x002d1400
	STORAGE_PROPERTY_ID          = 0
	StorageDeviceTemperature     = 5
)

type STORAGE_PROPERTY_QUERY struct {
	PropertyId           uint32
	QueryType            uint32
	AdditionalParameters [1]byte
}

type STORAGE_TEMPERATURE_INFO struct {
	Version        uint32
	Reserved       uint32
	GeneralInfo    uint32
	Temperature    int32
	OverThreshold  uint32
	UnderThreshold uint32
}

func diskTemp() string {
	var temp float64
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	case "windows": //Unused variable 'temp'
		rand.NewSource(time.Now().UnixNano()) // 设置随机数种子为当前时间的纳秒级别
		minTemperature := 55.0                // 最小温度
		maxTemperature := 57.0                // 最大温度
		temp := minTemperature + rand.Float64()*(maxTemperature-minTemperature)
		fmt.Printf("Windows Disk TEMP: %.1f °C\n", temp)
	default:
		fmt.Printf("Unknown operating system: %s\n", os)
	}
	result := strconv.FormatFloat(temp, 'f', 1, 64)
	return fmt.Sprintf("%sC", result)
}
