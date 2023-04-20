package main

import "fmt"

type PodTemplate interface {
	CreatePod()
	StoragePattern()
}

type Pod struct {
	PodName   string
	Namespace string
	Images    string
}

type BacupKubernetesSavePod struct {
	Pod
	EtcdAddress string
}

type BacupDevopsSavePod struct {
	Pod
	MinioAddress string
	MongoAddress string
	MysqlAddress string
}

type RestoreDevopsPod struct {
	Pod
}

// 创建pod
func (b *BacupKubernetesSavePod) CreatePod() {
	fmt.Println("创建pod")
}

func (b *BacupKubernetesSavePod) StoragePattern() {
	//TODO implement me
	fmt.Println("implement me")
}

func main() {
	var podTemplate PodTemplate
	podTemplate = &BacupKubernetesPod{}
	podTemplate.CreatePod()
}
