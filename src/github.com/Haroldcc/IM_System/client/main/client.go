package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	connect    net.Conn
	cmd        int // 客户端用户输入的菜单操作命令
}

// Client Creator
func NewClient(serverIp string, serverPort int) *Client {
	client := Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		cmd:        999,
	}

	// 连接server
	connect, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}

	client.connect = connect

	return &client
}

// 客户端菜单
func (client *Client) menu() bool {
	var choose int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&choose)

	if choose >= 0 && choose <= 3 {
		client.cmd = choose
		return true
	} else {
		fmt.Println(">>>>>>>>请输入合法选项<<<<<<<<<")
		return false
	}
}

// 更改用户名
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>>请输入用户名：")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.connect.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("Write error: ", err)
		return false
	}

	return true
}

// 处理server回应的消息（目前直接将消息打印到终端）
func (client *Client) DoResponse() {
	// 客户端连接一旦有消息，直接打印，（阻塞）
	io.Copy(os.Stdout, client.connect)
}

// 公聊模式
func (client *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>>>请输入聊天内容，exit退出.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.connect.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("Write error: ", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>请输入聊天内容，exit退出.")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) QueryOnlineUsers() {
	sendMsg := "who\n"
	_, err := client.connect.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("Write error: ", err)
		return
	}
}

// 私聊模式
func (client *Client) PrivateChat() {
	// 查询在线用户
	client.QueryOnlineUsers()
	fmt.Println(">>>>>请输入私聊对象[用户名],exit退出:")
	var remoteUserName string
	fmt.Scanln(&remoteUserName)

	for remoteUserName != "exit" {
		fmt.Println(">>>>>请输入消息内容，exit退出：")
		var chatMsg string
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteUserName + "|" + chatMsg + "\n\n"
				_, err := client.connect.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("Write error: ", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>>请输入消息内容，exit退出：")
			fmt.Scanln(&chatMsg)
		}

		remoteUserName = ""
		// 查询在线用户
		client.QueryOnlineUsers()
		fmt.Println(">>>>>请输入私聊对象[用户名],exit退出:")
		fmt.Scanln(&remoteUserName)
	}
}

// 启动客户端
func (client *Client) Run() {
	for client.cmd != 0 {
		for !client.menu() {
			// 循环显示菜单
		}

		// 根据不同命令处理处理相应业务
		switch client.cmd {
		case 1:
			fmt.Println("您已进入公聊模式...")
			client.PublicChat()
		case 2:
			fmt.Println("您已进入私聊模式...")
			client.PrivateChat()
		case 3:
			fmt.Println("您已进入更改用户名模式...")
			client.UpdateName()
		}
	}
}

var serverIp string // IP
var serverPort int  // 端口

// 通过命令行解析ip和端口
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8888）")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>连接服务器失败...")
		return
	}

	fmt.Println(">>>>>>>>连接服务器成功")

	// 开启goroutine，处理server的消息响应
	go client.DoResponse()

	// 启动客户端
	client.Run()
}
