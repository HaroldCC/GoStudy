/***********************************************************
 * 文件名称: IRequest.go
 * 功能描述: 请求抽象层
 * 创建标识: Haroldcc 2021/09/22
***********************************************************/

package ziface

// 请求接口：
// 将客户端的连接请求和请求数据，封装到Request中
type IRequest interface {
	/**
	 * @brief：获得当前连接
	 * @return 当前连接
	 */
	GetConnection() IConnection

	/**
	 * @brief：获得请求的数据
	 * @return 请求的数
	 */
	GetData() []byte
}
