/***********************************************************
 * 文件名称: router.go
 * 功能描述: 路由实现层
 * 创建标识: Haroldcc 2021/09/22
***********************************************************/

package znet

import "GoStudy/src/github.com/Haroldcc/zinx/ziface"

// 实现一个基类路由
// 可以重写IRouter中的所有接口，或继承BaseRouter，根据需要重写方法
type BaseRouter struct{}

/**
 * @brief：处理连接业务之前的方法
 * @param [in] request 请求
 */
func (router *BaseRouter) PrevHandle(request ziface.IRequest) {}

/**
 * @brief：处理连接业务的主方法
 * @param [in] request 请求
 */
func (router *BaseRouter) Handle(request ziface.IRequest) {}

/**
 * @brief：处理业务之后的方法
 * @param [in] request 请求
 */
func (router *BaseRouter) PostHandle(request ziface.IRequest) {}
