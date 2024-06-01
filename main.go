package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"monitor/hardware"
	"monitor/pkg"
	"monitor/proto"
	"runtime"
	"time"
)

const (
	LogLineRegex = `^\[(\d+/\d+/\d+\s+\d+:\d+:\d+)\]\s+(.*)$`
) // Arduino日志行的正则表达式
func main() {
	var path string
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		path = "/dev/ttyACM0"
	case "windows": //Unused variable 'temp'
		path = "COM14"
	default:
		fmt.Printf("Unknown operating system: %s\n", os)
	}
	// 打开串口
	c := &serial.Config{Name: path, Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	go serialListener(s)
	cpuMode := false // 用于切换 CPU 和网络模式
	count := 0
	ticker := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ticker:
			if cpuMode {
				cpuInfo2(err, s)
			} else {
				netInfo(err, s)
			}
			count++
			if count == 10 {
				cpuMode = !cpuMode // 切换模式
				count = 0
			}
		}
	}
}

func serialListener(s *serial.Port) {
	scanner := bufio.NewScanner(s)

	for scanner.Scan() {
		data := scanner.Bytes()
		// 判断数据包类型
		if len(data) > 0 && data[0] == proto.STX {
			// 自定义数据包
			command := data[1]
			dataLength := int(data[2])
			dataBlock := data[3 : 3+dataLength]
			receivedChecksum := data[3+dataLength]
			// 验证校验和
			expectedChecksum := proto.CalculateChecksum(dataBlock)
			pkg.Log.Printf(" %02X,%02X,%02X,%s,%02X", proto.STX, command, dataLength, string(dataBlock), receivedChecksum)
			if receivedChecksum == expectedChecksum {
				//pkg.Log.Printf("校验通过")
			} else {
				//fmt.Println("校验失败")
			}
		} else {
			// Arduino日志行
			logLine := string(data)
			pkg.Log.Printf("<-arduino %s", logLine)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func cpuInfo2(err error, s *serial.Port) {
	cpu := hardware.GetCPUInfo()
	bytes := proto.BuildMsg(proto.CmdCPU, cpu)
	_, err = s.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func netInfo(err error, s *serial.Port) {
	net, err := hardware.Net()
	bytes := proto.BuildMsg(proto.CmdNet, net)
	// 发送数据到串口
	_, err = s.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
