/***********************************************************
 * 文件名称: server.go
 * 功能描述: 服务端实现层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"errors"
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

		connId := uint32(0)
		// 3.阻塞等待客户端连接，处理客户端业务（读写）
		for {
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP error: ", err)
				continue
			}

			// 将当前连接业务处理与连接绑定
			handleConn := NewConnection(tcpConn, connId, EchoToClient)
			connId++

			// 启动业务处理
			go handleConn.Start()
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

/**
 * @brief：将内容回显给客户端
 * @param [in] conn:与客户端的连接
 * @param [in] data:数据
 * @param [in] count:数据的字节长度
 * @return 失败返回错误信息，成功返回nil
 */
func EchoToClient(conn *net.TCPConn, data []byte, count int) error {
	if _, err := conn.Write(data[:count]); err != nil {
		fmt.Println("Write back error: ", err)
		return errors.New("EchoToClient error")
	}

	return nil
}
