/***********************************************************
 * 文件名称: IConnManager.go
 * 功能描述: 连接管理抽象层
 * 创建标识: Haroldcc 2021/09/28
***********************************************************/

package ziface

// 连接管理抽象模块
type IConnManager interface {
	/**
	 * @brief：添加连接
	 * @param [in] conn 连接
	 */
	Add(conn IConnection)

	/**
	 * @brief：移除连接
	 * @param [in] conn 连接
	 */
	Remove(conn IConnection)

	/**
	 * @brief：根据ID获取连接
	 * @param [in] connID 连接ID
	 * @return 成功返回连接，错误返回错误信息
	 */
	Get(connID uint32) (IConnection, error)

	/**
	 * @brief：当前连接总数
	 * @return 连接总数
	 */
	Size() int

	/**
	 * @brief：清除并终止所有连接
	 */
	Clear()
}
