/***********************************************************
 * 文件名称: main.go
 * 功能描述: 服务器启动入口
 * 创建标识: Haroldcc 2021/10/11
***********************************************************/

package main

import (
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/apis"
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/core"
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"GoStudy/src/github.com/Haroldcc/zinx/znet"
	"fmt"
)

/**
 * @brief：客户端连接建立之后的Hook方法
 * @param [in] conn 连接
 */
func OnConnectionAdd(conn ziface.IConnection) {
	// 创建一个玩家
	player := core.NewPlayer(conn)

	// 给客户端发送【MsgID:1】消息:同步playerId
	player.SyncPlayerId()

	// 给客户端发送【MsgID:200】消息:同步player初始位置
	player.BroadCastBornPosition()

	// 将当前新上线的玩家添加到WordManager中
	core.WordManagerObj.AddPlayer(player)

	// 给当前连接绑定一个属性，记录当前连接的玩家ID
	conn.SetProperty("playerId", player.PlayerID)

	// 同步周边玩家，告知当前玩家（自己）已上线，广播当前玩家位置信息
	player.SyncNearbyPlayers()

	fmt.Println("===>Player id=", player.PlayerID, "is online<===")
}

func main() {
	// 创建server句柄
	server := znet.NewServer("MMO Game Server")

	// 注册连接建立后和连接销毁前的Hook方法
	server.SetOnConnStart(OnConnectionAdd)

	// 注册一些路由业务
	server.AddRouter(2, &apis.WorldChatApi{})
	server.AddRouter(3, &apis.PlayerMoveApi{})

	// 启动服务器
	server.Run()
}
