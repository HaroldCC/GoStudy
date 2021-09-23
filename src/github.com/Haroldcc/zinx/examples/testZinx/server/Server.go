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

// 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// test PreHandle
func (router *PingRouter) PrevHandle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call PreHandle error: ", err)
	}

}

// test Handle
func (router *PingRouter) Handle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping ping...\n"))
	if err != nil {
		fmt.Println("call Handle error: ", err)
	}

}

// test PostHandle
func (router *PingRouter) PostHandle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call PostHandle error: ", err)
	}

}

func main() {
	// 1.创建一个server
	server := znet.NewServer("[zinx服务端测试v0.4]")

	// 2.给框架添加一个自定义的router
	server.AddRouter(&PingRouter{})

	// 3.启动server
	server.Run()
}
