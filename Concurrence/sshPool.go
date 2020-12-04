package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"sync"
	"time"
)

type PoolConn struct {
	client   *ssh.Client
	mu       sync.RWMutex
	device   *Devices
	unusable bool
}

type Devices struct {
	sshHost     string
	sshUser     string
	sshPassword string
	sshPort     int
}

// 标记当前连接是否已断
func (p *PoolConn) sign() {
	p.mu.Lock()
	p.unusable = true
	p.mu.Unlock()
}

func (p *PoolConn) unsign() {
	p.mu.Lock()
	p.unusable = false
	p.mu.Unlock()
}

var ClientPool = make(map[string]*PoolConn, 150)

func PutSshPool(deviceID string, device *PoolConn) {
	ClientPool[deviceID] = device
}

func GetSshPool(key string) {
	var client = ClientPool[key]
	if client != nil {
		execute(client, "ls")
	}
}

func main() {
	var device = Devices{
		sshHost:     "172.168.1.76",
		sshUser:     "huoshen",
		sshPassword: "123456",
		sshPort:     22,
	}

	var poolConn = PoolConn{
		device:   &device,
		unusable: true,
	}

	var isCreate = CreateSshClient(&poolConn)
	if isCreate != nil {
		poolConn.client = isCreate
		PutSshPool("172.168.1.76", &poolConn)
		println("创建成功")
	}
	go Reconnection()

	GetSshPool("172.168.1.76")

	select {}
}

func CreateSshClient(conn *PoolConn) *ssh.Client {
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            conn.device.sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(conn.device.sshPassword)}

	addr := fmt.Sprintf("%s:%d", conn.device.sshHost, conn.device.sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)

	if err != nil {
		log.Fatal("创建ssh client 失败", err)
		return nil
	}
	return sshClient

}

func execute(conn *PoolConn, cmd string) {
	if conn != nil {
		conn.mu.Lock()
		defer conn.mu.Unlock()
		session, err := conn.client.NewSession()
		if err != nil {
			log.Fatal("创建ssh session 失败", err)
		}
		defer session.Close()
		//执行远程命令
		combo, err := session.CombinedOutput(cmd)
		if err != nil {
			conn.unusable = false
			log.Fatal("远程执行cmd 失败", err)
		}
		log.Println("命令输出:", string(combo))
	} else {
		println("sshClient 为空")
	}
}

func Reconnection() {
	for {
		for key, value := range ClientPool {
			if !value.unusable {
				var client = CreateSshClient(value)
				if client != nil {
					value.client = client
				} else {
					delete(ClientPool, key)
				}
			}
			println(key)
		}
		time.Sleep(time.Minute)
	}

}
