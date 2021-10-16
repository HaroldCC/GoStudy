/***********************************************************
 * 文件名称: world_chat.go
 * 功能描述: 世界聊天模块
 * 创建标识: Haroldcc 2021/10/15
***********************************************************/

package apis

import (
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/core"
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/pb"
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"GoStudy/src/github.com/Haroldcc/zinx/znet"
	"fmt"

	"github.com/golang/protobuf/proto"
)

// 世界聊天路由业务
type WorldChatApi struct {
	znet.BaseRouter
}

/**
 * @brief：处理连接业务的主方法
 * @param [in] request 请求
 */
func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	// 1 解析客户端传来的proto协议
	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Talk message UnMarshal error:", err)
		return
	}

	// 2 获取当前发送聊天数据的玩家
	playerId, err := request.GetConnection().GetProperty("playerId")
	if err != nil {
		fmt.Println("player id=", playerId, "not exist.")
		return
	}

	// 3 根据playerId获得对应的player
	player := core.WordManagerObj.GetPlayerById(playerId.(int32))

	// 4 将消息广播给其它全部在线玩家
	player.Talk(proto_msg.Content)
}
