/***********************************************************
 * 文件名称: IMessage.go
 * 功能描述: 消息模块抽象层
 * 创建标识: Haroldcc 2021/09/24
***********************************************************/

package ziface

// 消息抽象接口
type IMessage interface {
	/**
	 * @brief：获取消息ID
	 * @return 消息ID
	 */
	GetMsgID() uint32

	/**
	 * @brief：获取消息长度IMessage
	 * @return 消息长度
	 */
	GetMsgSize() uint32

	/**
	 * @brief：获取消息内容
	 * @return 消息内容
	 */
	GetMsgContent() []byte

	/**
	 * @brief：设置消息ID
	 * @param [in] id 消息ID
	 */
	SetMsgID(id uint32)

	/**
	 * @brief：设置消息长度IMessage
	 * @param [in] size 消息长度
	 */
	SetMsgSize(size uint32)

	/**
	 * @brief：设置消息内容
	 * @param [in] content 消息内容
	 */
	SetMsgContent(content []byte)
}
