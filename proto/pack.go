package proto

import (
	"fmt"
	"monitor/pkg"
)

// Pack struct
// | 起始位 | 命令位 | 数据长度位 | 数据块 | 校验位 |
//
// - **起始位**: 固定值，用于标识消息的开始。例如：0x02（STX - Start of Text）。
// - **命令位**: 标识消息类型的一个字节。例如：0x01 表示请求数据，0x02 表示发送数据，等等。
// - **数据长度位**: 标识数据块长度的一个字节。
// - **数据块**: 实际传输的数据，长度可变。
// - **校验位**: 校验和或其他校验机制，用于确保数据完整性。
type Pack struct {
	Header byte
	Cmd    byte
	Length uint16
	Data   []byte
	CRC    uint32
}

// BuildMsg 构建协议报文
func BuildMsg(command byte, data string) []byte {
	dataLength := len(data)
	message := []byte{StartByte, command, byte(dataLength)}
	message = append(message, []byte(data)...)
	checksum := calculateChecksum([]byte(data))
	message = append(message, checksum)
	pkg.Log.Printf(" %x %s ", command, data)
	pkg.Log.Printf(" %x %x %x %x %x ", StartByte, command, dataLength, data, checksum)
	return message
}

// ParseMsg 解析协议报文
func ParseMsg(message []byte) (byte, byte, string, byte, error) {
	if message[0] != StartByte {
		return 0, 0, "", 0, fmt.Errorf("invalid start byte")
	}

	command := message[1]
	dataLength := int(message[2])
	data := string(message[3 : 3+dataLength])
	checksum := message[3+dataLength]

	if calculateChecksum(message[:3+dataLength]) != checksum {
		return 0, 0, "", 0, fmt.Errorf("invalid checksum")
	}

	return command, byte(dataLength), data, checksum, nil
}

// 异或校验
func calculateChecksum(data []byte) byte {
	var checksum byte
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}
