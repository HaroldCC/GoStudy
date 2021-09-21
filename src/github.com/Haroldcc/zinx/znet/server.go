/***********************************************************
 * 文件名称: server.go
 * 功能描述: 服务端实现层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"fmt"
	"net"
)

// IServer的接口实现：定义一个Server的服务器模块
type Server struct {
	Name      string // 服务器名称
	IPVersion string // 服务器版本
	IP        string // 服务器IP
	Port      int    // 服务器端口
}

// 启动服务器
func (server *Server) Start() {
	fmt.Printf("[Start]Server Listener at IP:%s,Port:%d,is starting\n",
		server.IP, server.Port)

	go func() {
		// 1.获取一个TCP的Addr
		address, err := net.ResolveTCPAddr(server.IPVersion,
			fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			fmt.Println("resolve tcp address error: ", err)
			return
		}

		// 2.监听地址
		listener, err := net.ListenTCP(server.IPVersion, address)
		if err != nil {
			fmt.Println("listen ", server.IPVersion, " error: ", err)
			return
		}

		fmt.Println("Start zinx server success, ", server.Name, "begin to listening")

		// 3.阻塞等待客户端连接，处理客户端业务（读写）
		for {
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP error: ", err)
				continue
			}

			// 处理连接的客户端业务
			// (v0.1:最大512字节的回显业务)
			go func() {
				for {
					buf := make([]byte, 512)
					count, err := tcpConn.Read(buf)
					if err != nil {
						fmt.Println("Read error: ", err)
						continue
					}

					fmt.Printf("recv Client message:%s, count %d\n", buf, count)

					// 回显
					if _, err := tcpConn.Write(buf[:count]); err != nil {
						fmt.Println("Write back error: ", err)
						continue
					}
				}
			}()
		}
	}()
}

// 停止服务器
func (server *Server) Stop() {

}

// 运行服务器
func (server *Server) Run() {
	server.Start()

	// 阻塞
	select {}
}

// 初始化Server(创建一个Server)
func NewServer(name string) ziface.IServer {
	server := Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}

	return &server
}
