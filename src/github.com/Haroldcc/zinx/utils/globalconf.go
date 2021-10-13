/***********************************************************
 * 文件名称: globalconf.go
 * 功能描述: 全局配置模块
 * 创建标识: Haroldcc 2021/09/23
***********************************************************/
package utils

import (
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

// zinx框架的全局配置参数
// 使用方可通过zinx.json进行配置
type GlobalConf struct {
	/*Server*/
	TcpServer ziface.IServer // 当前zinx全局的server对象
	Host      string         // 当前服务器主机监听的IP
	TcpPort   int            // 当前服务器主机监听的端口
	Name      string         // 当前服务器名称

	/*zinx*/
	Version           string // zinx版本号
	MaxConn           int    // 服务器主机允许的最大连接数
	MaxPackageSize    uint32 // zinx框架数据包的最大值
	WorkerPoolSize    uint32 // 当前工作协程的数量
	MaxWorkerTaskSize uint32 // zinx框架限定的每个工作协程任务队列的最大任务数量
}

// 初始化GlobalConf
func init() {
	// 默认初始化
	G_config = &GlobalConf{
		Host:              "0.0.0.0",
		TcpPort:           8888,
		Name:              "zinxServer app",
		Version:           "v0.9",
		MaxConn:           1000,
		MaxPackageSize:    4096,
		WorkerPoolSize:    10,
		MaxWorkerTaskSize: 1024,
	}

	// 尝试加载使用方的配置参数
	G_config.LoadConfig()
}

/**
 * @brief：从文件中加载配置信息
 */
func (conf *GlobalConf) LoadConfig() {
	configData, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	// 将json数据解析到struct中
	err = json.Unmarshal(configData, &G_config)
	if err != nil {
		panic(err)
	}
}

// 全局的对外配置对象
var G_config *GlobalConf
