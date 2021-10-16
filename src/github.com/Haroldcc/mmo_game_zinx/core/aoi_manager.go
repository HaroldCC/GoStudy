/***********************************************************
 * 文件名称: aoi_manager.go
 * 功能描述: AOI区域管理模块
 * 创建标识: Haroldcc 2021/10/10
***********************************************************/

package core

import "fmt"

// 边界值
const (
	AOI_MIN_X   int = 85
	AOI_MAX_X   int = 410
	AOI_MIN_Y   int = 75
	AOI_MAX_Y   int = 400
	AOI_COUNT_X int = 10
	AOI_COUNT_Y int = 20
)

// AOI(Area Of Interest)区域管理器
type AOIManager struct {
	UpEdge     int           // 当前区域上边边界值
	DownEdge   int           // 当前区域下边边界值
	LeftEdge   int           // 当前区域左边边界值
	RightEdge  int           // 当前区域右边边界值
	GridCountX int           // 横向格子数量
	GridCountY int           // 纵向格子数量
	Grids      map[int]*Grid // 当前区域有哪些格子
}

/**
 * @brief：初始化AOI区域管理器
 * @param [in]
 * @param [out]
 * @return
 */
func NewAOIManager(upEdge, downEdge, leftEdge, rightEdge, gridCountX, gridCountY int) *AOIManager {
	aoiMgr := &AOIManager{
		UpEdge:     upEdge,
		DownEdge:   downEdge,
		LeftEdge:   leftEdge,
		RightEdge:  rightEdge,
		GridCountX: gridCountX,
		GridCountY: gridCountY,
		Grids:      make(map[int]*Grid),
	}

	// 对AOI初始化区域的所有格子进行编号和初始化
	for y := 0; y < gridCountY; y++ {
		for x := 0; x < gridCountX; x++ {
			// 格子编号：id = idY * gridCountX + idX
			gid := y*gridCountX + x

			// 初始化编号为gid的格子
			aoiMgr.Grids[gid] = NewGrid(gid,
				aoiMgr.LeftEdge+x*aoiMgr.gridWidthX(),
				aoiMgr.LeftEdge+(x+1)*aoiMgr.gridWidthX(),
				aoiMgr.UpEdge+y*aoiMgr.gridWidthY(),
				aoiMgr.UpEdge+(y+1)*aoiMgr.gridWidthY())
		}
	}

	return aoiMgr
}

/**
 * @brief：获得单个格子横向宽度
 * @return 横向宽度
 */
func (aoiMgr *AOIManager) gridWidthX() int {
	return (aoiMgr.RightEdge - aoiMgr.LeftEdge) / aoiMgr.GridCountX
}

/**
 * @brief：获得单个格子纵向宽度
 * @return 格子纵向宽度
 */
func (aoiMgr *AOIManager) gridWidthY() int {
	return (aoiMgr.DownEdge - aoiMgr.UpEdge) / aoiMgr.GridCountY
}

/**
 * @brief：输出AOI管理区域格子的基本信息
 * @return 格子基本信息
 */
func (aoiMgr *AOIManager) String() string {
	// AOIManager基本信息
	gridInfo := fmt.Sprintf(`AOIManager:\n leftEdge:%d,rightEdge:%d, 
	upEdge:%d,downEdge:%d,gridCountX:%d,gridWidthY:%d`,
		aoiMgr.LeftEdge, aoiMgr.RightEdge,
		aoiMgr.UpEdge, aoiMgr.RightEdge, aoiMgr.GridCountX, aoiMgr.GridCountY)

	// AOIManager管理区域的所有格子基本信息
	for _, grid := range aoiMgr.Grids {
		gridInfo += fmt.Sprintln(grid)
	}

	return gridInfo
}

/**
 * @brief：根据所在格子的GID得到周边九宫格格子集合
 * @param [in] gID 所处的格子ID
 * @return 周边格子集合
 */
func (aoiMgr *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	// 判断编号gID的格子是否存在
	if _, ok := aoiMgr.Grids[gID]; !ok {
		return
	}

	// 将当前编号gID的格子加入到九宫格切片中
	grids = append(grids, aoiMgr.Grids[gID])

	// 得到当前格子横向编号 indexX = gID % GridCountX
	indexX := gID % aoiMgr.GridCountX

	// 判断当前格子左边是否有格子
	if indexX > 0 {
		grids = append(grids, aoiMgr.Grids[gID-1])
	}

	// 判断当前格子右边是否有格子
	if indexX < aoiMgr.GridCountX-1 {
		grids = append(grids, aoiMgr.Grids[gID+1])
	}

	// 将当前记录到的横向格子全部取出，进行遍历，得到横向格子的ID集合
	gridsX := make([]int, 0, len(grids))
	for _, val := range grids {
		gridsX = append(gridsX, val.GID)
	}

	// 根据横向格子的ID集合进行遍历，得到横向格子的上下格子
	for _, val := range gridsX {
		// 得到当前格子纵向编号 indexY = gID / GridCountY
		indexY := val / aoiMgr.GridCountY

		// 上边是否有格子
		if indexY > 0 {
			grids = append(grids, aoiMgr.Grids[val-aoiMgr.GridCountX])
		}

		// 下边是否有格子
		if indexY < aoiMgr.GridCountY-1 {
			grids = append(grids, aoiMgr.Grids[val+aoiMgr.GridCountX])
		}
	}

	return grids
}

/**
 * @brief：根据横纵坐标得到当前所处的格子ID
 * @param [in] x 横坐标
 * @param [in] y 纵坐标
 * @return 当前所处的格子ID
 */
func (aoiMgr *AOIManager) GetGidByPos(x, y float32) int {
	/**************************************************************************
		注：设括号为一个格子，括号中的数值即为格子ID,横向为x轴，纵向为y轴，原点位于左上角。
			以下为一个5x5的网格示例

		(00) (01) (02) (03) (04)
		(05) (06) (07) (08) (09)
		(10) (11) (12) (13) (14)
		(15) (16) (17) (18) (19)
		(20) (21) (22) (23) (24)

	**************************************************************************/
	indexX := (int(x) - aoiMgr.UpEdge) / aoiMgr.gridWidthX()
	indexY := (int(y) - aoiMgr.LeftEdge) / aoiMgr.gridWidthY()

	return indexY*aoiMgr.GridCountX + indexX
}

/**
 * @brief：根据横纵坐标得到周边九宫格内全部的PlayerIDs
 * @param [in] x 横坐标
 * @param [in] y 纵坐标
 * @return 周边九宫格内的PlayerID集合
 */
func (aoiMgr *AOIManager) GetPlayerIdsByPos(x, y float32) (playerIDs []int) {
	// 得到当前坐标的格子ID
	gID := aoiMgr.GetGidByPos(x, y)

	// 得到周边九宫格
	grids := aoiMgr.GetSurroundGridsByGid(gID)

	// 得到将周边九宫格中的玩家ID
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		//fmt.Printf("==>gird id:%d,playerIDs:%v==\n", grid.GID, grid.GetPlayerIDs())
	}

	return playerIDs
}

/**
 * @brief：添加一个playerID到一个格子中
 * @param [in] playerID 玩家ID
 * @param [out] gID 格子ID
 */
func (aoiMgr *AOIManager) AddPlayerIdToGridByGid(playerID, gID int) {
	aoiMgr.Grids[gID].AddPlayer(playerID)
}

/**
 * @brief：从一个格子中移除一个playerID
 * @param [in] playerID 玩家ID
 * @param [out] gID 格子ID
 */
func (aoiMgr *AOIManager) RemovePlayerIdFromGridByGid(playerID, gID int) {
	aoiMgr.Grids[gID].RemovePlayer(playerID)
}

/**
 * @brief：通过一个格子的gID获取全部的playerID
 * @param [in] playerID 玩家ID
 * @param [out] gID 格子ID
 * return playerID集合
 */
func (aoiMgr *AOIManager) GetPlayerIdsByGid(playerID, gID int) (playerIDs []int) {
	playerIDs = aoiMgr.Grids[gID].GetPlayerIDs()
	return
}

/**
 * @brief：通过坐标将playerID添加到格子中
 * @param [in] playerID 玩家ID
 * @param [out] gID 格子ID
 * return playerID集合
 */
func (aoiMgr *AOIManager) AddPlayerIdToGridByPos(playerID int, x, y float32) {
	aoiMgr.AddPlayerIdToGridByGid(playerID, aoiMgr.GetGidByPos(x, y))
}

/**
 * @brief：通过坐标将playerID从格子中移除
 * @param [in] playerID 玩家ID
 * @param [out] gID 格子ID
 * return playerID集合
 */
func (aoiMgr *AOIManager) RemovePlayerIdFromGridByPos(playerID int, x, y float32) {
	aoiMgr.RemovePlayerIdFromGridByGid(playerID, aoiMgr.GetGidByPos(x, y))
}
