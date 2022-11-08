package main

import (
	firstpb "LearningNotes-Go/example/grpc-demo/firstpb"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
)

func main() {
	pm := NewPersonMessage()
	//if err := writeToFile("person.bin", pm); err != nil {
	//	log.Fatalln("写入失败")
	//}

	//pm2 := &firstpb.PersonMessage{}
	//_ = readFromFile("person.bin", pm2)

	pmStr := toJson(pm)
	fmt.Println(pmStr)

	pm3 := &firstpb.PersonMessage{}
	_ = fromJson(pmStr, pm3)
	fmt.Println("pb struct :", pm3)
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("转化为Json时发生错误", err.Error())
	}

	return str
}

func fromJson(in string, pb proto.Message) error {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("读取Json时发生错误", err.Error())
	}
	return nil
}

func NewPersonMessage() *firstpb.PersonMessage {
	pm := firstpb.PersonMessage{
		Id:           1234,
		IsAdult:      true,
		Name:         "test",
		LuckyNumbers: []int32{1, 2, 3, 4, 5},
	}

	pm.Name = "Nick"
	return &pm
}

func writeToFile(fileName string, pb proto.Message) error {
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("无法序列化到 bytes")
	}

	if err := os.WriteFile(fileName, dataBytes, 0644); err != nil {
		log.Fatalln("无法写入到文件")
	}

	log.Println("成功写入文件")

	return nil
}

func readFromFile(fileName string, pb proto.Message) error {
	dataBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("读取文件时发生错误")
	}

	err = proto.Unmarshal(dataBytes, pb)
	if err != nil {
		log.Fatalln("转换错误")
	}

	return nil
}
