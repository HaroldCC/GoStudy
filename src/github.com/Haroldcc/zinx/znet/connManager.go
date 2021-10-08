/***********************************************************
 * 文件名称: connManager.go
 * 功能描述: 连接管理模块实现层
 * 创建标识: Haroldcc 2021/09/29
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

// 连接管理模块
type ConnManager struct {
	connections map[uint32]ziface.IConnection // 管理的连接集合
	connMapLock sync.RWMutex                  // 连接集合读写锁
}

/**
 * @brief：创建连接管理
 * @param [in]
 * @param [out]
 * @return
 */
func NewConnMgr() ziface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

/**
 * @brief：添加连接
 * @param [in] conn 连接
 */
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	// 加写锁
	connMgr.connMapLock.Lock()
	defer connMgr.connMapLock.Unlock()

	// 将conn添加到ConnManager中
	connMgr.connections[conn.GetConnectionId()] = conn
	fmt.Println("[connection id=", conn.GetConnectionId(),
		"add to ConnManager succeed.Connection sum=", connMgr.Size(), "]")
}

/**
* @brief：移除连接
* @param [in] conn 连接
 */
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	// 加写锁
	connMgr.connMapLock.Lock()
	defer connMgr.connMapLock.Unlock()

	// 将conn从Connmanager中移除
	delete(connMgr.connections, conn.GetConnectionId())
	fmt.Println("[connection id=", conn.GetConnectionId(),
		"remove from ConnManager succeed.Connection sum=", connMgr.Size(), "]")
}

/**
 * @brief：根据ID获取连接
 * @param [in] connID 连接ID
 * @return 成功返回连接，错误返回错误信息
 */
func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	// 加读锁
	connMgr.connMapLock.RLock()
	defer connMgr.connMapLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

/**
 * @brief：当前连接总数
 * @return 连接总数
 */
func (connMgr *ConnManager) Size() int {
	return len(connMgr.connections)
}

/**
 * @brief：清除并终止所有连接
 */
func (connMgr *ConnManager) Clear() {
	// 加写锁
	connMgr.connMapLock.Lock()
	defer connMgr.connMapLock.Unlock()

	// 删除conn,并停止conn的工作
	for connID, conn := range connMgr.connections {
		// 停止
		conn.Stop()

		// 删除
		delete(connMgr.connections, connID)
	}

	fmt.Println("[Clear all connections succeed.Connection sum=", connMgr.Size())
}
