/***********************************************************
 * 文件名称: request.go
 * 功能描述: 请求实现层
 * 创建标识: Haroldcc 2021/09/22
***********************************************************/

package znet

import "GoStudy/src/github.com/Haroldcc/zinx/ziface"

type Request struct {
	conn ziface.IConnection // 与客户端建立好的连接
	msg  ziface.IMessage    // 客户端请求的数据
}

/**
 * @brief：获得当前连接
 * @return 当前连接
 */
func (request *Request) GetConnection() ziface.IConnection {
	return request.conn
}

/**
 * @brief：获得请求的数据
 * @return 请求的数
 */
func (request *Request) GetData() []byte {
	return request.msg.GetMsgContent()
}

/**
 * @brief：获取消息ID
 * @return 消息ID
 */
func (request *Request) GetMsgID() uint32 {
	return request.GetMsgID()
}
