/***********************************************************
 * 文件名称: main.go
 * 功能描述: 服务器启动入口
 * 创建标识: Haroldcc 2021/10/11
***********************************************************/

package main

import "GoStudy/src/github.com/Haroldcc/zinx/znet"

func main() {
	// 创建server句柄
	server := znet.NewServer("MMO Game Server")

	// 注册连接建立后和连接销毁前的Hook方法

	// 注册一些路由业务

	// 启动服务器
	server.Start()
}
