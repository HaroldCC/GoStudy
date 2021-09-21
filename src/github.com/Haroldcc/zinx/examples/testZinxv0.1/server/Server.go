/***********************************************************
 * 文件名称: Server.go
 * 功能描述: zinx服务端测试
 * 创建标识: Haroldcc 2021/09/21
***********************************************************/

package main

import "GoStudy/src/github.com/Haroldcc/zinx/znet"

func main() {
	server := znet.NewServer("[zinx服务端测试v0.1]")
	server.Run()
}
