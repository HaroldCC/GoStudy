/***********************************************************
 * 文件名称: datapack_test.go
 * 功能描述: 数据包模块单元测试
 * 创建标识: Haroldcc 2021/09/25
***********************************************************/

package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(test *testing.T) {
	/* 模拟服务器 */

	// 创建一个Tcp监听
	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("server listen err: ", err)
		return
	}

	// 创建一个协程，负责模拟客户端处理业务
	go func() {
		for {
			// 接受客户端连接
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server Accept error: ", err)
			}

			go func(conn net.Conn) {
				// 处理客户端请求（拆包）
				dp := NewDataPack()
				for {
					// 1.从包中读出包头
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error: ", err)
						break
					}

					// 将包头解包，读出里面的DataLen
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("server unpack error: ", err)
						return
					}

					if msgHead.GetMsgSize() > 0 {
						msg := msgHead.(*Message) // 类型断言
						msg.content = make([]byte, msg.GetMsgSize())

						// 2.根据包头存储的长度，读出数据内容
						_, err := io.ReadFull(conn, msg.content)
						if err != nil {
							fmt.Println("server unpack data error: ", err)
							return
						}

						// 将读到的消息内容显示
						fmt.Println(">>>>>>>>>Recv MsgID: ", msg.id,
							"dataSize: ", msg.size,
							"content: ", string(msg.content))
					}
				}
			}(conn)
		}
	}()

	/* 模拟客户端 */
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("client dial error: ", err)
		return
	}

	db := NewDataPack()

	// 创建消息
	msg1 := Message{
		id:      1,
		size:    4,
		content: []byte("zinx"),
	}

	sendData1, err := db.Pack(&msg1)
	if err != nil {
		fmt.Println("client pack message error: ", err)
		return
	}

	msg2 := Message{
		id:      2,
		size:    5,
		content: []byte("Hello"),
	}
	sendData2, err := db.Pack(&msg2)
	if err != nil {
		fmt.Println("client pack message error: ", err)
		return
	}

	// 模拟粘包，将两条message一同发送
	sendData1 = append(sendData1, sendData2...)

	// 发送数据
	conn.Write(sendData1)

	// 客户端阻塞
	select {}
}
