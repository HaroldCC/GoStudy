/***********************************************************
 * 文件名称: server.go
 * 功能描述: 服务端实现层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/utils"
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"fmt"
	"net"
)

// IServer的接口实现：定义一个Server的服务器模块
type Server struct {
	Name        string                        // 服务器名称
	IPVersion   string                        // 服务器版本
	IP          string                        // 服务器IP
	Port        int                           // 服务器端口
	MsgHandler  ziface.IMsgHandle             // 消息管理模块，绑定MsgID和对应的业务处理API
	ConnMgr     ziface.IConnManager           // 连接管理模块
	OnConnStart func(conn ziface.IConnection) // server创建连接之后调用的Hook方法
	OnConnStop  func(conn ziface.IConnection) // server关闭连接之前调用的Hook方法
}

// 启动服务器
func (server *Server) Start() {
	fmt.Printf("[zinx server Name:%s, listener at IP:%s, port:%d is starting]\n",
		utils.G_config.Name, utils.G_config.Host, utils.G_config.TcpPort)
	fmt.Printf("[zinx version %s, MaxConn:%d, MaxPackageSize:%d]\n",
		utils.G_config.Version, utils.G_config.MaxConn, utils.G_config.MaxPackageSize)

	go func() {
		// 0.开启worker工作池
		server.MsgHandler.StartWorkPool()

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

			// 判断当前连接个数，超过配置中最大连接，关闭新的连接
			if server.ConnMgr.Size() >= utils.G_config.MaxConn {
				fmt.Println("[Too many connections, MaxConn=", utils.G_config.MaxConn, "]")
				tcpConn.Close()
				continue
			}

			// 将当前连接业务处理与连接绑定
			handleConn := NewConnection(server, tcpConn, connId, server.MsgHandler)
			connId++

			// 启动业务处理
			go handleConn.Start()
		}
	}()
}

// 停止服务器
func (server *Server) Stop() {
	fmt.Println("[Stop!zinx server name: ", server.Name)
	server.ConnMgr.Clear()
}

// 运行服务器
func (server *Server) Run() {
	server.Start()

	// 阻塞
	select {}
}

/**
 * @brief：路由功能:给当前的服务注册一个路由方法，共客户端连接处理
 * @param [in] msgID 消息ID
 * @param [in] router 路由
 */
func (server *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	server.MsgHandler.AddRouter(msgID, router)
}

/**
 * @brief：获取连接管理器
 * @return 连接管理器
 */
func (server *Server) GetConnMgr() ziface.IConnManager {
	return server.ConnMgr
}

// 初始化Server(创建一个Server)
func NewServer(name string) ziface.IServer {
	server := Server{
		Name:       utils.G_config.Name,
		IPVersion:  "tcp4",
		IP:         utils.G_config.Host,
		Port:       utils.G_config.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnMgr(),
	}

	return &server
}

/**
 * @brief：注册连接创建之后的钩子方法
 * @param [in] hookFunc Hook方法
 */
func (server *Server) SetOnConnStart(hookFunc func(conn ziface.IConnection)) {
	server.OnConnStart = hookFunc
}

/**
 * @brief：注册连接销毁之前的钩子方法
 * @param [in] hookFunc Hook方法
 */
func (server *Server) SetOnConnStop(hookFunc func(conn ziface.IConnection)) {
	server.OnConnStop = hookFunc
}

/**
 * @brief：调用连接创建之后的钩子方法
 * @param [in] conn 连接
 */
func (server *Server) CallOnConnStart(conn ziface.IConnection) {
	if server.OnConnStart != nil {
		fmt.Println("[****Call OnConnStart()...")
		server.OnConnStart(conn)
	}
}

/**
 * @brief：调用连接销毁之前的钩子方法
 * @param [in] conn 连接
 */
func (server *Server) CallOnConnStop(conn ziface.IConnection) {
	if server.OnConnStop != nil {
		fmt.Println("[****Call OnConnStop()...]")
		server.OnConnStop(conn)
	}
}
