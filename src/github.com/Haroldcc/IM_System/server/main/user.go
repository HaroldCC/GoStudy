package main

import (
	"net"
	"strings"
)

// 用户
//每个User维持一个与客户端的连接，用以与客户端的通信
type User struct {
	Name    string
	Address string
	channel chan string // 与服务器通信的管道
	connect net.Conn    // 与客户端的连接
	server  *Server     // 当前用户所在的server
}

// User Creator
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Address: userAddr,
		channel: make(chan string),
		connect: conn,
		server:  server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 监听当前User的channel，一旦有消息，直接发送给客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.channel

		user.connect.Write([]byte((msg + "\n")))
	}
}

// 用户上线
func (user *User) Online() {
	// 用户上线，将用户加入OnlineMap
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 广播用户上线消息
	user.server.BoardCast(user, "已上线")
}

// 用户下线
func (user *User) Offline() {
	// 用户下线，将用户从OnlineMap中删掉
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	// 广播用户下线消息
	user.server.BoardCast(user, "已下线")
}

// 用户消息处理
func (user *User) HandleMessage(msg string) {
	if msg == "who" {
		// 用户查询当前在线用户
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineMsg := "[" + u.Address + "]" + u.Name + ":" + "在线...\n"
			user.SendMsgToClient(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 用户改名（消息格式：rename|张三）
		newName := msg[7:]

		// 判断newName是否被其它用户占用
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsgToClient("当前用户名已被占用\n")
		} else {
			// 更名(先删除，再添加)
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.Name = newName
			user.SendMsgToClient("您已更新用户名：" + user.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私聊功能（消息格式：to|张三|消息内容）
		// 获取对方用户名
		remoteUserName := strings.Split(msg, "|")[1]
		if remoteUserName == "" {
			user.SendMsgToClient("消息格式不正确，Usage：to|张三|hello")
			return
		}

		// 根据用户名，查询到对方user对象
		remoteUser, ok := user.server.OnlineMap[remoteUserName]
		if !ok {
			user.SendMsgToClient("未查询到用户：" + remoteUserName)
			return
		}

		// 获取到消息内容，通过对方的user对象发送消息内容
		content := strings.Split(msg, "|")[2]
		remoteUser.SendMsgToClient(user.Name + "对您说：" + content)
	} else {
		// 用户消息处理(目前是直接将消息全局广播)
		user.server.BoardCast(user, msg)
	}
}

// 给当前user对应的客户端发送消息
func (user *User) SendMsgToClient(msg string) {
	user.connect.Write([]byte(msg))
}
