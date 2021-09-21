/***********************************************************
 * 文件名称: connection.go
 * 功能描述: 连接实现层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"fmt"
	"net"
)

// IConnection模块实现
type Connection struct {
	Conn     *net.TCPConn   // 当前连接的TCP socket
	ConnID   uint32         // 连接的ID
	isClosed bool           // 连接状态
	handler  ziface.Handler // 当前连接绑定的业务处理handler
	ExitChan chan bool      // 告知当前连接状态的channel(退出/停止)
}

/**
 * @brief：创建一个Connection
 * @param [in] conn: 连接的socket
 * @param [in] connID:连接ID
 * @param [in] callback:处理业务回调函数
 * @return 创建的connection
 */
func NewConnection(conn *net.TCPConn, connID uint32, callback ziface.Handler) *Connection {
	connection := Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		handler:  callback,
		ExitChan: make(chan bool, 1),
	}

	return &connection
}

/**
 * @brief：处理读业务
 */
func (conn *Connection) StartReader() {
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println("ConnID=", conn.ConnID, " Reader is exit, remote address is ",
		conn.Conn.RemoteAddr().String())
	defer conn.Stop()

	for {
		buf := make([]byte, 512)
		count, err := conn.Conn.Read(buf)
		if err != nil {
			fmt.Println("Read error ", err)
			continue
		}

		// 调用当前连接绑定的业务处理方法
		if err := conn.handler(conn.Conn, buf, count); err != nil {
			fmt.Println("ConnID=", conn.ConnID, " handle is error", err)
			break
		}
	}

}

/**
 * @brief：启动连接，开始工作
 */
func (conn *Connection) Start() {
	fmt.Println("Conn Start...ConnID=", conn.ConnID)

	// 启动读业务
	go conn.StartReader()
}

/**
 * @brief：停止连接，结束工作
 */
func (conn *Connection) Stop() {
	fmt.Println("Conn Stop...ConnID= ", conn.ConnID)

	// 当前连接是否关闭
	if conn.isClosed {
		return
	}

	// 关闭连接
	conn.Conn.Close()

	// 回收资源
	close(conn.ExitChan)

}

/**
 * @brief：获取当前连接的绑定socket connection
 * @return 当前绑定的连接
 */
func (conn *Connection) GetTcpConnection() *net.TCPConn {
	return conn.Conn
}

/**
 * @brief：获取当前连接的ID
 * @return 连接ID
 */
func (conn *Connection) GetConnectionId() uint32 {
	return conn.ConnID
}

/**
 * @brief：获取远程客户端的TCP状态（IP+Port)
 * @return 客户端地址
 */
func (conn *Connection) GetRemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

/**
 * @brief：发送数据，将数据发送至远程客户端
 * @param [in] data:发送的数据
 * @return 失败返回错误信息，成功返回nil
 */
func (conn *Connection) Send(data []byte) error {
	return nil
}
