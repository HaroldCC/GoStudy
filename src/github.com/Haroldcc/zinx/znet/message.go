/***********************************************************
 * 文件名称: message.go
 * 功能描述: 消息模块实现层
 * 创建标识: Haroldcc 2021/09/24
***********************************************************/

package znet

type Message struct {
	id      uint32 // 消息ID
	size    uint32 // 消息长度
	content []byte // 消息内容
}

/**
 * @brief：获取消息ID
 * @return 消息ID
 */
func (message *Message) GetMsgID() uint32 {
	return message.id
}

/**
 * @brief：获取消息长度IMessage
 * @return 消息长度
 */
func (message *Message) GetMsgSize() uint32 {
	return message.size
}

/**
 * @brief：获取消息内容
 * @return 消息内容
 */
func (message *Message) GetMsgContent() []byte {
	return message.content
}

/**
 * @brief：设置消息ID
 * @param [in] id 消息ID
 */
func (message *Message) SetMsgID(id uint32) {
	message.id = id
}

/**
 * @brief：设置消息长度IMessage
 * @param [in] size 消息长度
 */
func (message *Message) SetMsgSize(size uint32) {
	message.size = size
}

/**
 * @brief：设置消息内容
 * @param [in] content 消息内容
 */
func (message *Message) SetMsgContent(content []byte) {
	message.content = content
}
