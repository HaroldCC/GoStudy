/***********************************************************
 * 文件名称: msgHandler.go
 * 功能描述: 消息管理实现层
 * 创建标识: Haroldcc 2021/09/25
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"fmt"
)

// 消息管理模块
type MsgHandle struct {
	APIs map[uint32]ziface.IRouter // 消息ID与其对应的处理接口map
}

/**
 * @brief：创建一个消息管理
 * @return 消息管理
 */
func NewMsgHandle() ziface.IMsgHandle {
	return &MsgHandle{
		APIs: make(map[uint32]ziface.IRouter),
	}
}

/**
 * @brief：执行对应的Router消息处理
 * @param [in] request 请求
 */
func (msgHandle *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	// 1.从request中找到msgID
	handler, ok := msgHandle.APIs[request.GetMsgID()]
	if !ok {
		fmt.Println("msgID=", request.GetMsgID(), "binging api is not found! needed to register!")
	}

	// 2.根据msgID调度对应的Router业务
	handler.PrevHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

/**
 * @brief：添加消息处理
 * @param [in] msgID 消息ID
 * @param [in] router 路由
 */
func (msgHandle *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	// 判断当前msgID绑定的API是否已存在
	if _, ok := msgHandle.APIs[msgID]; ok {
		// msgID绑定的API已注册
		panic(fmt.Sprintf("repeat api, msgID= %d", msgID))
	}

	// 添加msgID与API的绑定
	msgHandle.APIs[msgID] = router

	fmt.Println("add api msgID= ", msgID, "succeed.")
}
