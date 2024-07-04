package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听8080端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	for {
		// 接受新的连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept failed: %v\n", err)
			continue
		}

		// 处理连接
		go process(conn)
	}
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()

	for {
		// 首先，创建一个缓冲区来读取3字节
		buf := make([]byte, 3)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("get Len failed: %v\n", err)
			continue
		}

		// 将第3字节转换为整数，作为帧的长度
		frameLen := uint16(buf[2])
		//获取帧全部信息
		buf = make([]byte, frameLen)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Printf("Read failed: %v\n", err)
			continue
		}

		checksum := buf[len(buf)-5]                       // 获取校验和
		calculated := calculateChecksum(buf[:len(buf)-5]) // 计算校验和
		if checksum != calculated {
			continue
		}

		//帧类型
		frameType := uint8(buf[5])
		switch frameType {
		case 0x00:
		case 0x01:
		case 0x02:
		case 0x03:
		case 0x04:
		case 0x05:
		case 0x06:
		case 0x07:
		case 0x08:
		case 0x09:
		case 0x0A:
		case 0x0B:
		case 0xF0:
		case 0xF1:
		case 0xF2:
		case 0xF3:
		case 0xF4:
		case 0xF5:
		case 0xF6:
		case 0xF7:
		default:
			fmt.Printf("Unknown frametype: %v\n", frameType)
		}
	}
}

// 计算校验和函数
func calculateChecksum(data []byte) byte {
	var sum int
	for _, b := range data {
		sum += int(b)
	}
	return byte(sum & 0xFF) // 取低字节
}
