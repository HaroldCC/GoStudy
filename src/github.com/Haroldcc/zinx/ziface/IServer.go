/***********************************************************
 * 文件名称: IServer.go
 * 功能描述: 服务端抽象层
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package ziface

// 服务器接口
type IServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 运行服务器
	Run()

	/**
	 * @brief：路由功能:给当前的服务注册一个路由方法，共客户端连接处理
	 * @param [in] msgID 消息ID
	 * @param [in] router 路由
	 */
	AddRouter(msgID uint32, router IRouter)
}
