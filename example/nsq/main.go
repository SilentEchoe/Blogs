package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var (
	tcpNsqdAddrr = "127.0.0.1:4150"
)

// 声明一个结构体，实现HandleMessage接口方法
type NsqHandler struct {
	//消息数
	msqCount int64
	//标识ID
	nsqHandlerID string
}

// 实现HandleMessage方法
// message是接收到的消息
func (s *NsqHandler) HandleMessage(message *nsq.Message) error {
	//没收到一条消息+1
	s.msqCount++
	fmt.Println(s.msqCount, s.nsqHandlerID)
	fmt.Printf("msg.Timestamp=%v, msg.nsqaddress=%s,msg.body=%s \n", time.Unix(0, message.Timestamp).Format("2006-01-02 03:04:05"), message.NSQDAddress, string(message.Body))
	return nil
}

func main() {
	config := nsq.NewConfig()
	//创造消费者，参数一是订阅的主题，参数二是使用的通道
	com, err := nsq.NewConsumer("Insert", "channel", config)
	if err != nil {
		fmt.Println(err)
	}
	//添加处理回调
	com.AddHandler(&NsqHandler{nsqHandlerID: "One"})
	//连接对应的nsqd
	err = com.ConnectToNSQD(tcpNsqdAddrr)
	if err != nil {
		fmt.Println(err)
	}
}
