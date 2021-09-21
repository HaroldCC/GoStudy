package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip          string           // ip
	Port        int              // 端口
	OnlineMap   map[string]*User // 在线用户列表
	mapLock     sync.RWMutex
	MessageChan chan string // 消息广播channel
}

// Server Creator
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:          ip,
		Port:        port,
		OnlineMap:   make(map[string]*User),
		MessageChan: make(chan string),
	}
	return server
}

// Server Starter
func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}

	// close socket (注：Go语言defer关键字特性)
	defer listener.Close()

	// 启动监听MessageChan的goroutine
	go server.ListenMessager()

	for {
		// accept connect
		connect, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			continue
		}

		// do handler
		go server.Handler(connect)
	}
}

// Server Handler
func (server *Server) Handler(connect net.Conn) {
	user := NewUser(connect, server)

	// 用户上线
	user.Online()

	// 监听用户是否活跃的channel
	isAlive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			len, err := connect.Read(buf)
			if len == 0 {
				// 用户下线
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read error: ", err)
				return
			}

			// 提取用户的消息（裁剪掉'\n'）
			msg := string(buf[:len-1])

			// 用户消息处理
			user.HandleMessage(msg)

			// 设置用户活跃
			isAlive <- true
		}
	}()

	// 判断用户是否活跃
	for {
		select {
		case <-isAlive:
			// 用户活跃，重置定时器
			// 不做任何事，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 300):
			// 超时强踢
			user.SendMsgToClient("您已被踢出")

			// 销毁资源
			close(user.channel)

			// 关闭连接
			connect.Close()

			// 退出Handler
			return // or runtime.Goexit()
		}
	}
}

// 广播消息
func (server *Server) BoardCast(user *User, msg string) {
	sendMsg := "[" + user.Address + "]" + user.Name + ": " + msg
	server.MessageChan <- sendMsg
}

// 监听MessageChan的goroutine，一旦有消息就发送给全部的在线user
func (server *Server) ListenMessager() {
	for {
		msg := <-server.MessageChan

		// 将msg发送给全部的在线user
		server.mapLock.Lock()
		for _, user := range server.OnlineMap {
			user.channel <- msg
		}
		server.mapLock.Unlock()
	}
}
