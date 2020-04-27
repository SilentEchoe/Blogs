package main

import (
	"fmt"
	firstpb "github.com/AnAnonymousFriend/LearningNotes-Go/src/first"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main()  {
	pm := NewPersonMessage()
	writeToFile("person.bin",pm)
}

func writeToFile(fileName string, pb proto.Message) error  {
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("无法序列化")
	}

	if err := ioutil.WriteFile(fileName,dataBytes, 0644);
	err != nil {
		log.Fatalln("无法写入到文件", err.Error())
	}
log.Println("成功")

return nil
}


func NewPersonMessage() *firstpb.PersonMessage {
	pm := firstpb.PersonMessage{
		Id:          1234,
		IsAdult:     true,
		Name:        "Dave",
		LuckNumvers: []int32{1,2,3,4,5},

	}

	fmt.Println(pm)

	pm.Name="Nick"

	fmt.Println(pm)

	fmt.Printf("The Id is %d\n",pm.GetId())

	return  &pm
}



