package hardware

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

func Net() (string, error) {
	prevNetStat, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("获取网络信息时出错: %s", err)
		return "", err
	}

	time.Sleep(time.Second) // 等待一秒钟

	currNetStat, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("获取网络信息时出错: %s", err)
		return "", nil
	}

	// 计算每秒的进站和出站数据
	incomingKbps := float64(currNetStat[0].BytesRecv-prevNetStat[0].BytesRecv) / 1024
	outgoingKbps := float64(currNetStat[0].BytesSent-prevNetStat[0].BytesSent) / 1024

	fmt.Printf("每秒进站数据: %.2f kbps, 每秒出站数据: %.2f kbps\n", incomingKbps, outgoingKbps)
	return fmt.Sprintf("%.2fkbps,%.2fkbps", incomingKbps, outgoingKbps), nil
}
