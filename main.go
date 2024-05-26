package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"monitor/hardware"
	"monitor/proto"
	"strings"
	"time"
)

func main() {

	// 打开串口
	c := &serial.Config{Name: "COM14", Baud: 115200} // 根据实际情况修改串口名称
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	// 用于存储未完成的消息
	var buffer strings.Builder

	// 启动一个 Goroutine 用于监听串口返回的数据
	go func() {
		buf := make([]byte, 128)
		for {
			n, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			buffer.Write(buf[:n])

			// 处理完整的消息
			for {
				msg := buffer.String()
				index := strings.Index(msg, "\n")
				if index == -1 {
					break
				}

				completeMessage := msg[:index]
				fmt.Printf("Received: %s\n", completeMessage)
				buffer.Reset()
				buffer.WriteString(msg[index+1:])
			}
		}
	}()

	for {
		//根据当前状态采集数据并发送
		cpuInfo(err, s)

		// 等待1秒钟
		time.Sleep(1 * time.Second)
	}
}

func cpuInfo(err error, s *serial.Port) {
	cpu := hardware.GetCPUInfo()
	bytes := proto.BuildMsg(proto.CmdCPU, cpu)

	// 创建新的字节切片，长度比原始字节切片长1
	newBytes := make([]byte, len(bytes)+1)

	// 复制原始字节到新的字节切片中
	copy(newBytes, bytes)

	// 在新的字节切片末尾添加换行符
	newBytes[len(newBytes)-1] = '\n'

	// 发送数据到串口
	_, err = s.Write(newBytes)
	if err != nil {
		log.Fatal(err)
	}
}
