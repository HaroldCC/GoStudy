/***********************************************************
 * 文件名称: player_move.go
 * 功能描述: 玩家移动模块
 * 创建标识: Haroldcc 2021/10/16
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

// 玩家移动
type PlayerMoveApi struct {
	znet.BaseRouter
}

/**
 * @brief：处理连接业务的主方法
 * @param [in] request 请求
 */
func (m *PlayerMoveApi) Handle(request ziface.IRequest) {
	// 解析客户端传递的proto协议数据
	proto_msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Player move: player position Unmarshal error:", err)
		return
	}

	// 得到当前发送位置的玩家
	playerId, err := request.GetConnection().GetProperty("playerId")
	if err != nil {
		fmt.Println("GetProperty playerId error:", err)
		return
	}

	fmt.Printf("Player id=%d,move(%f,%f,%f,%f)\n", playerId,
		proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

	// 广播当前玩家的位置给其他玩家
	player := core.WordManagerObj.GetPlayerById(playerId.(int32))
	player.BroadCastPosition(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)
}
