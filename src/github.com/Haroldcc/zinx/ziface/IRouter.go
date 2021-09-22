/***********************************************************
 * 文件名称: IRouter.go
 * 功能描述: 路由抽象层
 * 创建标识: Haroldcc 2021/09/22
***********************************************************/

package ziface

// 路由抽象接口
// 路由里的数据都是IRequest
type IRouter interface {
	/**
	 * @brief：处理连接业务之前的方法
	 * @param [in] request 请求
	 */
	PrevHandle(request IRequest)

	/**
	 * @brief：处理连接业务的主方法
	 * @param [in] request 请求
	 */
	Handle(request IRequest)

	/**
	 * @brief：处理业务之后的方法
	 * @param [in] request 请求
	 */
	PostHandle(request IRequest)
}
