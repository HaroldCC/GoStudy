/***********************************************************
 * 文件名称: IConnection.go
 * 功能描述: 连接抽象层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package ziface

import "net"

// 连接模块接口
type IConnection interface {
	/**
	 * @brief：启动连接，开始工作
	 */
	Start()

	/**
	 * @brief：停止连接，结束工作
	 */
	Stop()

	/**
	 * @brief：获取当前连接的绑定socket connection
	 * @return 当前绑定的连接
	 */
	GetTcpConnection() *net.TCPConn

	/**
	 * @brief：获取当前连接的ID
	 * @return 连接ID
	 */
	GetConnectionId() uint32

	/**
	 * @brief：获取远程客户端的TCP状态（IP+Port)
	 * @return 客户端地址
	 */
	GetRemoteAddr() net.Addr

	/**
	 * @brief：发送数据，将数据发送至远程客户端
	 * @param [in] data:发送的数据
	 * @return 失败返回错误信息，成功返回nil
	 */
	Send(data []byte) error
}

// 业务处理方法
type Handler func(*net.TCPConn, []byte, int) error
