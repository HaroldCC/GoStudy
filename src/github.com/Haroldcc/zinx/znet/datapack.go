/***********************************************************
 * 文件名称: datapack.go
 * 功能描述: 数据包实现层
 * 创建标识: Haroldcc 2021/09/24
***********************************************************/

package znet

import (
	"GoStudy/src/github.com/Haroldcc/zinx/utils"
	"GoStudy/src/github.com/Haroldcc/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

// IDataPack模块实现
type DataPack struct {
}

/**
 * @brief：获取包头长度
 * @return 包头长度
 */
func (dataPace *DataPack) GetHeadLen() uint32 {
	// DataLen:uint32(4字节) + MsgID:uint32(4字节)
	return 8
}

/**
 * @brief：封包 包体格式：|DataLen|MsgID|Content|
 * @param [in] msg 消息体
 * @return 成功返回封装的二进制消息流，失败error!=nil
 */
func (dataPace *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建存放字节流的buffer
	dataBuffer := bytes.NewBuffer([]byte{})

	// 1.写入DataLen
	if err := binary.Write(dataBuffer,
		binary.LittleEndian,
		msg.GetMsgSize()); err != nil {
		return nil, err
	}

	// 2.写入MsgID
	if err := binary.Write(dataBuffer,
		binary.LittleEndian,
		msg.GetMsgID()); err != nil {
		return nil, err
	}

	// 3.写入Content
	if err := binary.Write(dataBuffer,
		binary.LittleEndian,
		msg.GetMsgContent()); err != nil {
		return nil, err
	}

	return dataBuffer.Bytes(), nil
}

/**
 * @brief：拆包 包体格式：|DataLen|MsgID|Content|
 * @param [in] data 数据流
 * @return 成功返回消息体，失败error!=nil
 */
func (dataPace *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	// 创建存放数据的buffer
	dataBuffer := bytes.NewReader(data)

	msg := Message{}

	// 1.读取DataLen
	if err := binary.Read(dataBuffer,
		binary.LittleEndian,
		&msg.size); err != nil {
		return nil, err
	}

	// 2.读取MsgID
	if err := binary.Read(dataBuffer,
		binary.LittleEndian,
		&msg.id); err != nil {
		return nil, err
	}

	// 判断DataLen是否超出允许的最大长度
	if utils.G_config.MaxPackageSize > 0 && msg.size > utils.G_config.MaxPackageSize {
		return nil, errors.New("too large message data recv")
	}

	return &msg, nil
}

/**
 * @brief：创建一个数据包
 * @param [in]
 * @param [out]
 * @return
 */
func NewDataPack() ziface.IDataPack {
	dataPack := DataPack{}
	return &dataPack
}
