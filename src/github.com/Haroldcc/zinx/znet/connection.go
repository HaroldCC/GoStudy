/***********************************************************
 * 文件名称: connection.go
 * 功能描述: 连接实现层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

// IConnection模块实现
type Connection struct {
	Conn       *net.TCPConn      // 当前连接的TCP socket
	ConnID     uint32            // 连接的ID
	isClosed   bool              // 连接状态
	ExitChan   chan bool         // 告知当前连接状态的channel(退出/停止)
	MsgHandler ziface.IMsgHandle // 消息管理模块，绑定MsgID和对应的业务处理API
}

/**
 * @brief：创建一个Connection
 * @param [in] conn: 连接的socket
 * @param [in] connID:连接ID
 * @param [in] msgHandler:消息处理API
 * @return 创建的connection
 */
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) ziface.IConnection {
	connection := Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
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
		// 创建数据包
		dataPackage := NewDataPack()

		// 读取包头
		dataHead := make([]byte, dataPackage.GetHeadLen())
		if _, err := io.ReadFull(conn.GetTcpConnection(),
			dataHead); err != nil {
			fmt.Println("read message head error: ", err)
			break
		}

		// 对包头拆包，存入message中
		msg, err := dataPackage.UnPack(dataHead)
		if err != nil {
			fmt.Println("unpack dataHead error: ", err)
			break
		}

		// 根据包头信息读取数据
		var data []byte
		if msg.GetMsgSize() > 0 {
			data = make([]byte, msg.GetMsgSize())
			if _, err := io.ReadFull(conn.GetTcpConnection(),
				data); err != nil {
				fmt.Println("read message data error: ", err)
				break
			}
		}

		//存入消息内容
		msg.SetMsgContent(data)

		// 得到conn的Request
		req := Request{
			conn: conn,
			msg:  msg,
		}

		// 根据绑定好的msgID,找到对应的API处理业务
		go conn.MsgHandler.DoMsgHandler(&req)
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
 * @param [in] msgID:消息ID
 * @param [in] data:发送的数据
 * @return 失败返回错误信息，成功返回nil
 */
func (conn *Connection) SendMsg(msgID uint32, data []byte) error {
	if conn.isClosed {
		return errors.New("Connection closed when send message")
	}

	// 对发送的数据进行封包
	dataPackage := NewDataPack()
	binaryMsg, err := dataPackage.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println("Pack error MsgID = ", msgID)
		return errors.New("Pack error message")
	}

	// 将数据包发送给客户端
	if _, err := conn.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write message id = ", msgID, " error: ", err)
		return errors.New("conn Write error")
	}

	return nil
}
