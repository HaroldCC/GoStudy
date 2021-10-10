/***********************************************************
 * 文件名称: AOIMgr_test.go
 * 功能描述: AOI模块单元测试
 * 创建标识: Haroldcc 2021/10/10
***********************************************************/

package tests

import (
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/core"
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoiMgr := core.NewAOIManager(0, 250, 0, 250, 5, 5)

	fmt.Println(aoiMgr)
}

func TestGetSurroundGridsByGid(t *testing.T) {
	aoiMgr := core.NewAOIManager(0, 250, 0, 250, 5, 5)

	for gid := range aoiMgr.Grids {
		grids := aoiMgr.GetSurroundGridsByGid(gid)
		fmt.Println("gid:", gid, "grids len=", len(grids))
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}

		fmt.Println("surrounding grid IDs are", gIDs)
	}
}
