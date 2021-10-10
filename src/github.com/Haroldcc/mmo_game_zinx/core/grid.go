/***********************************************************
 * 文件名称: grid.go
 * 功能描述: AOI模型的格子
 * 创建标识: Haroldcc 2021/10/10
***********************************************************/

package core

import (
	"fmt"
	"sync"
)

// AOI模型的格子类型
type Grid struct {
	GID           int          // 网格id
	UpEdge        int          // 当前格子上边边界值
	DownEdge      int          // 当前格子下边边界值
	LeftEdge      int          // 当前格子左边边界值
	RightEdge     int          // 当前格子右边边界值
	playerIDs     map[int]bool // 当前格子内的玩家或物体的ID集合
	playerIDsLock sync.RWMutex // 锁
}

/**
 * @brief：初始化当前格子
 * @param [in] gID 格子id
 * @param [in] upEdge 上边界
 * @param [in] downEdge 下边界
 * @param [in] leftEdge 左边界
 * @param [in] rightEdge 右边界
 * @return 初始化好的的格子
 */
func NewGrid(gID, upEdge, downEdge, leftEdge, rightEdge int) *Grid {
	return &Grid{
		GID:       gID,
		UpEdge:    upEdge,
		DownEdge:  downEdge,
		LeftEdge:  leftEdge,
		RightEdge: rightEdge,
		playerIDs: make(map[int]bool),
	}
}

/**
 * @brief：添加玩家
 * @param [in] playerID 玩家id
 */
func (grid *Grid) AddPlayer(playerID int) {
	grid.playerIDsLock.Lock()
	grid.playerIDs[playerID] = true
	grid.playerIDsLock.Unlock()
}

/**
 * @brief：移除玩家
 * @param [in] playerID 玩家id
 */
func (grid *Grid) RemovePlayer(playerID int) {
	grid.playerIDsLock.Lock()
	delete(grid.playerIDs, playerID)
	grid.playerIDsLock.Unlock()
}

/**
 * @brief：获得当前格子中的所有玩家id
 * @return 当前格子中的玩家id集合
 */
func (grid *Grid) GetPlayerIDs() (playerIDs []int) {
	grid.playerIDsLock.RLock()
	defer grid.playerIDsLock.RUnlock()

	for key := range grid.playerIDs {
		playerIDs = append(playerIDs, key)
	}

	return playerIDs
}

/**
 * @brief：输出格子的基本信息
 * @return 格子的基本信息
 */
func (grid *Grid) String() string {
	return fmt.Sprintf("Grid id:%d,upEdge:%d,downEdge:%d,leftEdge:%d,rightEdge:%d,playerIDs:%v,",
		grid.GID, grid.UpEdge, grid.DownEdge, grid.LeftEdge, grid.RightEdge, grid.playerIDs)
}
