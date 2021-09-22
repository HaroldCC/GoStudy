/***********************************************************
 * 文件名称: Client.go
 * 功能描述: 服务端
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Start...")

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Client start error, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("Hello zinx v0.3..."))
		if err != nil {
			fmt.Println("Write conn error: ", err)
			return
		}

		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read buf error: ", err)
			return
		}

		fmt.Printf("Server send back:%s, count=%d\n", buf, count)

		// cpu阻塞1秒
		time.Sleep(1 * time.Second)
	}
}
