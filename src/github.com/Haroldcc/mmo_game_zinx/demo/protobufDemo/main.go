package main

import (
	"GoStudy/src/github.com/Haroldcc/mmo_game_zinx/demo/protobufDemo/pb"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := pb.Person{
		Name:  "zhangsan",
		Email: "zhangsan@qq.com",
		Phones: []*pb.Person_PhoneNumber{
			&pb.Person_PhoneNumber{
				Number: "11111111",
				Type:   pb.Person_MOBILE,
			},
			&pb.Person_PhoneNumber{
				Number: "2222222",
				Type:   pb.Person_HOME,
			},
			&pb.Person_PhoneNumber{
				Number: "3333333",
				Type:   pb.Person_WORK,
			},
		},
	}

	/*编码*/
	// 将person对象，就是将protobuf的message进行序列化
	data, err := proto.Marshal(&person)
	// data就是进行网络传输的数据，对端按照message Person的格式进行发序列化
	if err != nil {
		fmt.Println("Marshal error:", err)
	}

	/*解码*/
	newPerson := pb.Person{}
	err = proto.Unmarshal(data, &newPerson)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
	}

	fmt.Println("源数据：", person)
	fmt.Println("解码之后的数据：", newPerson)
}
