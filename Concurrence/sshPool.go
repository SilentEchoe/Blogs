package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
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

func GetShhClient(clineName string) *net.Conn {
	return nil
}

var ClientPool = make(map[string]*PoolConn, 150)

func PutSshPool(deviceID string, device *PoolConn) {
	println(device.device.sshHost)
	println(device.device.sshUser)
	ClientPool[deviceID] = device
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
		PutSshPool("172.168.1.76", &poolConn)
	}
	println(*isCreate)

}

func SSH() {
	sshHost := "172.168.1.76"
	sshUser := "huoshen"
	sshPassword := "123456"
	sshPort := 22

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout: time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:    sshUser,
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)

	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	//执行远程命令
	combo, err := session.CombinedOutput("whoami; cd /; ls -al;echo https://github.com/dejavuzhou/felix")
	if err != nil {
		log.Fatal("远程执行cmd 失败", err)
	}
	log.Println("命令输出:", string(combo))
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

func GetSshClient() *ssh.Client {
	return nil
}
