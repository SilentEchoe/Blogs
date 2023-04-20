package main

import "fmt"

const (
	SaveToMinio = 1
	SaveToSftp  = 2
	SaveToLocal = 3
)

type Storage struct {
	MinioEndpoint string
	MinioBucket   string
	MinioFile     string

	SftpEndpoint string
	SftpFilePath string

	VolueFile string
}

type PodTemplate interface {
	CreatePod()
	StoragePattern(storeType int)
}

type Pod struct {
	PodName   string
	Namespace string
	Images    string
	Args      []string
}

type BacupKubernetesSavePod struct {
	Pod
	BaseStorage Storage
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

//CreatePod 创建 BacupKubernetesSavePod
func (b *BacupKubernetesSavePod) CreatePod() {
	newArgs := []string{"cluster-etcd-endpoint=", b.EtcdAddress}

	b.Args = append(b.Args, newArgs...)
	fmt.Println("创建 BacupKubernetesSavePod 结构体", b.Args)
}

func (b *BacupKubernetesSavePod) StoragePattern(storeType int) {
	switch storeType {
	case SaveToMinio:
		newArgs := []string{"store-minio-endpoint=", b.BaseStorage.MinioEndpoint, "store-minio-bucket=" + b.BaseStorage.MinioBucket, "store-minio-file=" + b.BaseStorage.MinioFile}
		b.Args = append(b.Args, newArgs...)
	case SaveToSftp:
		newArgs := []string{"store-sftp-endpoint=", b.BaseStorage.SftpEndpoint, "store-sftp-file=" + b.BaseStorage.SftpFilePath}
		b.Args = append(b.Args, newArgs...)
	case SaveToLocal:
		newArgs := []string{"store-volume-file=", b.BaseStorage.VolueFile}
		b.Args = append(b.Args, newArgs...)
	}

	fmt.Println("设置 BacupKubernetesSavePod的存储方式", b.Args)
}

func main() {
	var podTemplate PodTemplate
	podTemplate = &BacupKubernetesSavePod{EtcdAddress: "https://127.0.0.1:2179", BaseStorage: Storage{VolueFile: "/tmp/backup.tar.gz"}}
	podTemplate.CreatePod()
	podTemplate.StoragePattern(SaveToLocal)
}
