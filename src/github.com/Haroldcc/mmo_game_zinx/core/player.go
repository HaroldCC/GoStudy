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
		PlayerID:    player.PlayerID,
		MessageType: 2,
		Data: &pb.BroadCast_Pos{
			Pos: &pb.Position{X: player.X, Y: player.Y, Z: player.Z, V: player.V},
		},
	}

	player.SendMsg(200, proto_msg)
}

/**
 * @brief：玩家广播世界聊天消息
 * @param [in] content 消息内容
 */
func (player *Player) Talk(content string) {
	// 1 组建MsgID:200的proto数据
	proto_msg := &pb.BroadCast{
		PlayerID:    player.PlayerID,
		MessageType: 1,
		Data:        &pb.BroadCast_Content{Content: content},
	}

	// 2 得到当前世界所有在线玩家
	players := WordManagerObj.GetAllPlayers()

	// 3 向所有玩家（包括自己）发送MsgID:200的消息
	for _, p := range players {
		p.SendMsg(200, proto_msg)
	}
}

/**
 * @brief：获取当前玩家周围玩家
 * @return 周围玩家集合
 */
func (player *Player) GetNearbyPlayers() []*Player {
	playerIds := WordManagerObj.AoiMgr.GetPlayerIdsByPos(player.X, player.Z)
	players := make([]*Player, 0, len(playerIds))
	for _, playerId := range playerIds {
		players = append(players, WordManagerObj.GetPlayerById(int32(playerId)))
	}

	return players
}

/**
 * @brief：同步附近玩家
 */
func (player *Player) SyncNearbyPlayers() {
	// 1 获取当前玩家周围的玩家
	players := player.GetNearbyPlayers()

	// 2 将当前玩家的位置信息发送给周围玩家（让其他玩家看到自己）
	// 2.1 组建MsgID:200 proto数据
	proto_msg := &pb.BroadCast{
		PlayerID:    player.PlayerID,
		MessageType: 2,
		Data: &pb.BroadCast_Pos{
			Pos: &pb.Position{X: player.X, Y: player.Y, Z: player.Z, V: player.V},
		},
	}
	// 2.2 将当前玩家位置广播给周围玩家
	for _, nearbyPlayer := range players {
		nearbyPlayer.SendMsg(200, proto_msg)
	}

	// 3 将周围玩家的信息发给当前玩家（让自己看到其他玩家）
	// 3.1 组件MsgID:202 proto数据
	// 3.1.1 制作pb.Player切片,存储周围玩家
	players_proto_msg := make([]*pb.Player, 0, len(players))
	for _, nearbyPlayer := range players {
		p := &pb.Player{
			PlayerID:  nearbyPlayer.PlayerID,
			PlayerPos: &pb.Position{X: nearbyPlayer.X, Y: nearbyPlayer.Y, Z: nearbyPlayer.Z, V: nearbyPlayer.V},
		}

		players_proto_msg = append(players_proto_msg, p)
	}

	// 3.1.2 将pb.Player切片封装进protobuf数据
	syncPlayer_proto_msg := &pb.SyncPlayers{
		Player: players_proto_msg[:],
	}

	// 3.1.3 将周围玩家广播给当前玩家
	player.SendMsg(202, syncPlayer_proto_msg)
}

/**
 * @brief：更新并广播玩家坐标
 * @param [in] x 横坐标
 * @param [in] y 高度
 * @param [in] z 纵坐标
 * @param [in] v 旋转角度
 * @return
 */
func (player *Player) BroadCastPosition(x, y, z, v float32) {
	// 更新玩家坐标
	player.X = x
	player.Y = y
	player.Z = z
	player.V = v

	// 组建proto消息 MsgID:200 MessageType:4
	proto_msg := &pb.BroadCast{
		PlayerID:    player.PlayerID,
		MessageType: 4,
		Data:        &pb.BroadCast_Pos{Pos: &pb.Position{X: x, Y: y, Z: z, V: v}},
	}

	// 获取周围玩家
	players := player.GetNearbyPlayers()
	for _, nearbyPlayer := range players {
		// 将玩家当前位置信息广播给周围玩家
		nearbyPlayer.SendMsg(200, proto_msg)
	}
}

/**
 * @brief：玩家下线
 */
func (player *Player) Offline() {
	// 获得附近玩家
	players := player.GetNearbyPlayers()

	// 给周围玩家广播下线信息MsgID:201
	proto_msg := &pb.SyncPlayerID{
		PlayerID: player.PlayerID,
	}

	for _, nearbyPlayer := range players {
		nearbyPlayer.SendMsg(201, proto_msg)
	}

	// 移除玩家
	WordManagerObj.RemovePlayerById(player.PlayerID)
}
