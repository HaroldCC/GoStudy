/***********************************************************
 * 文件名称: IDataPack.go
 * 功能描述: 数据包抽象模块
 * 创建标识: Haroldcc 2021/09/24
***********************************************************/

package ziface

// 数据包模块
// 将TCP中的数据流封装为一个个的数据包，处理TCP粘包问题
type IDataPack interface {
	/**
	 * @brief：获取包头长度
	 * @return 包头长度
	 */
	GetHeadLen() uint32

	/**
	 * @brief：封包
	 * @param [in] msg 消息体
	 * @return 成功返回封装的二进制消息流，失败error!=nil
	 */
	Pack(msg IMessage) ([]byte, error)

	/**
	 * @brief：拆包
	 * @param [in] data 数据流
	 * @return 成功返回消息体，失败error!=nil
	 */
	UnPack(data []byte) (IMessage, error)
}
