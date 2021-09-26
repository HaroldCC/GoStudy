/***********************************************************
 * 文件名称: Server.go
 * 功能描述: zinx服务端测试
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package main

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"GoStudy/src/github.com/Haroldcc/zinx/znet"
	"fmt"
)

// 自定义路由1
type PingRouter struct {
	znet.BaseRouter
}

// test Handle
func (router *PingRouter) Handle(request ziface.IRequest) {
	// 读取客户端数据并打印
	fmt.Println("recv from client msgID=", request.GetMsgID(),
		", content:", string(request.GetData()))

	// 向客户端发送消息
	err := request.GetConnection().SendMsg(200, []byte("a message [Ping] form server"))
	if err != nil {
		fmt.Println("send message error: ", err)
	}
}

// 自定义路由2
type HelloRouter struct {
	znet.BaseRouter
}

// test Handle
func (router *HelloRouter) Handle(request ziface.IRequest) {
	// 读取客户端数据并打印
	fmt.Println("recv from client msgID=", request.GetMsgID(),
		", content:", string(request.GetData()))

	// 向客户端发送消息
	err := request.GetConnection().SendMsg(201, []byte("a message [HelloRouter] form server"))
	if err != nil {
		fmt.Println("send message error: ", err)
	}
}

func main() {
	// 1.创建一个server
	server := znet.NewServer("[zinx服务端测试v0.7]")

	// 2.给框架添加一个自定义的router1
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})

	// 3.启动server
	server.Run()
}
