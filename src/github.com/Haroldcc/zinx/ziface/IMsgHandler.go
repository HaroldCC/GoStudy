/***********************************************************
 * 文件名称: IMsgHandler.go
 * 功能描述: 消息管理模块抽象层
 * 创建标识: Haroldcc 2021/09/25
***********************************************************/

package ziface

// 消息管理接口
type IMsgHandle interface {
	/**
	 * @brief：执行对应的Router消息处理
	 * @param [in] request 请求
	 */
	DoMsgHandler(request IRequest)

	/**
	 * @brief：添加消息处理
	 * @param [in] msgID 消息ID
	 * @param [in] router 路由
	 */
	AddRouter(msgID uint32, router IRouter)
}
