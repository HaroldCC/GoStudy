/***********************************************************
 * 文件名称: Client.go
 * 功能描述: 服务端
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package main

import (
	"GoStudy/src/github.com/Haroldcc/zinx/znet"
	"fmt"
	"io"
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
		// 对消息进行封包
		dataPackage := znet.NewDataPack()
		binaryMsg, err := dataPackage.Pack(znet.NewMessage(0, []byte("zinx v0.5 client test message")))
		if err != nil {
			fmt.Println("Pack message error: ", err)
			return
		}

		// 向服务端发送消息包
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("Write data error: ", err)
			return
		}

		// 接收服务端发来的数据
		binaryHead := make([]byte, dataPackage.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error: ", err)
			break
		}

		msg, err := dataPackage.UnPack(binaryHead)
		if err != nil {
			fmt.Println("client unpack message head error: ", err)
			break
		}

		if msg.GetMsgSize() > 0 {
			// message有数据,读出数据
			data := make([]byte, msg.GetMsgSize())
			if _, err := io.ReadFull(conn, data); err != nil {
				fmt.Println("read message data error: ", err)
				return
			}

			// 设置消息内容
			msg.SetMsgContent(data)

			// 打印消息内容
			fmt.Println("----recv server message, msgID=", msg.GetMsgID(),
				" size=", msg.GetMsgSize(),
				" content:", string(msg.GetMsgContent()))
		}

		// cpu阻塞1秒
		time.Sleep(1 * time.Second)
	}
}
