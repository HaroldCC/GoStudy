/***********************************************************
 * 文件名称: word_manager.go
 * 功能描述: 世界管理模块
 * 创建标识: Haroldcc 2021/10/13
***********************************************************/

package core

import "sync"

// 游戏世界管理器
type WordManager struct {
	AoiMgr      *AOIManager       // 当前世界的AOI管理模块
	Players     map[int32]*Player // 当前世界全部在线玩家
	playersLock sync.RWMutex      // 锁
}

// 世界管理器句柄（全局）
var WordManagerObj *WordManager

// 初始化
func init() {
	WordManagerObj = &WordManager{
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_MIN_Y, AOI_MAX_Y, AOI_COUNT_X, AOI_COUNT_Y),
		Players: make(map[int32]*Player),
	}
}

/**
 * @brief：添加玩家
 * @param [in] plyer 玩家
 */
func (wm *WordManager) AddPlayer(player *Player) {
	wm.playersLock.Lock()
	wm.Players[player.PlayerID] = player
	wm.playersLock.Unlock()

	// 将player添加到AoiMgr中
	wm.AoiMgr.AddPlayerIdToGridByPos(int(player.PlayerID), player.X, player.Z)
}

/**
 * @brief：通过ID删除玩家
 * @param [in] playerID 玩家ID
 */
func (wm *WordManager) RemovePlayerById(playerID int32) {
	// 通过ID得到当前玩家
	player := wm.Players[playerID]

	// 将player从AoiMgr中删除
	wm.AoiMgr.RemovePlayerIdFromGridByPos(int(playerID), player.X, player.Z)

	wm.playersLock.Lock()
	delete(wm.Players, playerID)
	wm.playersLock.Unlock()
}

/**
 * @brief：通过ID获取玩家
 * @param [in] playerID 玩家ID
 * @return 玩家
 */
func (wm *WordManager) GetPlayerById(playerID int32) *Player {
	wm.playersLock.RLock()
	defer wm.playersLock.RUnlock()
	return wm.Players[playerID]
}

/**
 * @brief：获取全部在线玩家
 * @return 在线玩家集合
 */
func (wm *WordManager) GetAllPlayers() []*Player {
	wm.playersLock.RLock()
	defer wm.playersLock.RUnlock()

	players := make([]*Player, 0)

	for _, p := range wm.Players {
		players = append(players, p)
	}

	return players
}
