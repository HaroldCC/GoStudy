/***********************************************************
 * 文件名称: msgHandler.go
 * 功能描述: 消息管理实现层
 * 创建标识: Haroldcc 2021/09/25
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/utils"
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"fmt"
)

// 消息管理模块
type MsgHandle struct {
	APIs           map[uint32]ziface.IRouter // 消息ID与其对应的处理接口map
	TaskQueue      []chan ziface.IRequest    // worker(工作协程)任务队列
	WorkerPoolSize uint32                    // worker(工作协程)个数
}

/**
 * @brief：创建一个消息管理
 * @return 消息管理
 */
func NewMsgHandle() ziface.IMsgHandle {
	return &MsgHandle{
		APIs:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.G_config.WorkerPoolSize),
		WorkerPoolSize: utils.G_config.WorkerPoolSize,
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
		return
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

/**
 * @brief：启动worker工作池
 */
func (msgHandle *MsgHandle) StartWorkPool() {
	// 根据WorkerPoolSize分别开启worker，每个worker使用一个goroutine承载
	for i := 0; i < int(msgHandle.WorkerPoolSize); i++ {
		msgHandle.TaskQueue[i] = make(chan ziface.IRequest, utils.G_config.MaxWorkerTaskSize)

		go msgHandle.StartWorker(i, msgHandle.TaskQueue[i])
	}
}

/**
 * @brief：启动一个worker工作流程
 * @param [in] workerID 工作协程ID
 * @param [in] taskQueue 任务队列
 */
func (msgHandle *MsgHandle) StartWorker(workerID int, taskQueue chan ziface.IRequest) {
	// 循环阻塞等待对应的消息队列的消息
	for {
		msgHandle.DoMsgHandler(<-taskQueue)
	}
}

/**
 * @brief：将消息分配给任务队列，由worker进行处理
 * @param [in] request 消息请求
 */
func (msgHandle *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 将消息平均分配给不同的worker
	workID := request.GetConnection().GetConnectionId() % msgHandle.WorkerPoolSize
	fmt.Println("[Add ConnId=", request.GetConnection().GetConnectionId(),
		" request MsgID=", request.GetMsgID(),
		" to workID= ", workID, "]")

	// 将消息发送给对应的worker所在的TaskQueue
	msgHandle.TaskQueue[workID] <- request
}
