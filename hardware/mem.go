package hardware

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

func Mem() (string, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("获取内存信息时出错: %s", err)
		return "获取内存错误", err
	}

	totalGB := float64(vmStat.Total) / 1024 / 1024 / 1024
	usedGB := float64(vmStat.Used) / 1024 / 1024 / 1024

	fmt.Printf("内存总量: %.2f GB, 内存使用量: %.2f GB\n", totalGB, usedGB)
	return fmt.Sprintf("内存总量: %.2f GB, 内存使用量: %.2f GB", totalGB, usedGB), nil
}
