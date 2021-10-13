/***********************************************************
 * 文件名称: player.go
 * 功能描述: 玩家对象
 * 创建标识: Haroldcc 2021/10/12
***********************************************************/

package core

import (
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/pb"
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"fmt"
	"math/rand"
	"sync"

	"github.com/golang/protobuf/proto"
)

// 玩家对象
type Player struct {
	PlayerID int32              // 玩家ID
	Conn     ziface.IConnection // 当前玩家的连接（用于和客户端的连接）
	X        float32            // 平面X坐标
	Y        float32            // 高度
	Z        float32            // 平面Y坐标
	V        float32            // 旋转的0-360角度
}

var PlayerIDGen int32 = 1      // 用来生成玩家ID的计数器
var PlayerIDGenLock sync.Mutex // 保护PlayerIDGen的Mutex

/**
 * @brief：创建玩家
 * @param [in] conn 客户端连接标识
 * @return 玩家
 */
func NewPlayer(conn ziface.IConnection) *Player {
	// 生成一个玩家ID
	PlayerIDGenLock.Lock()
	id := PlayerIDGen
	PlayerIDGen++
	PlayerIDGenLock.Unlock()

	return &Player{
		PlayerID: id,
		Conn:     conn,
		X:        float32(160 + rand.Intn(10)), // 在160坐标点，基于X轴随机偏移若干
		Y:        0,
		Z:        float32(140 + rand.Intn(20)), // 在140坐标点，基于Y轴随机偏移若干
		V:        0,
	}
}

/**
 * @brief：发送消息
 * @param [in] msgId 消息ID
 * @param [in] data 发送的消息数据
 */
func (player *Player) SendMsg(msgId uint32, data proto.Message) {
	// 将proto Message结构体序列化成二进制数据
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Marshal message error:", err)
		return
	}

	if player.Conn == nil {
		fmt.Println("Connection in player is nil")
		return
	}

	// 发送二进制数据
	if err := player.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("Player send message error:", err)
		return
	}
}

/**
 * @brief：同步玩家ID至客户端
 */
func (player *Player) SyncPlayerId() {
	// 给客户端发送【MsgID:1】消息
	proto_msg := &pb.SyncPlayerID{
		PlayerID: player.PlayerID,
	}

	player.SendMsg(1, proto_msg)
}

/**
 * @brief：广播玩家的出生地点
 */
func (player *Player) BroadCastBornPosition() {
	// 给客户端发送【MsgID:200】消息
	proto_msg := &pb.BroadCast{
		PlayerID: player.PlayerID,
		Tp:       2,
		Data: &pb.BroadCast_Pos{
			Pos: &pb.Position{X: player.X, Y: player.Y, Z: player.Z, V: player.V},
		},
	}

	player.SendMsg(200, proto_msg)
}
