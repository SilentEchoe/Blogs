package main

import (
	"bytes"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/klog/v2"
	"text/template"
)

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

	VolumeMounts map[string]string
	Volumes      map[string]string

	IsHostPath bool
}

type PodTemplate interface {
	CreatePod()
	StoragePattern(storeType int)
}

type Pod struct {
	AppName   string
	Namespace string
	ImageName string
	Args      []string
}

type BacupKubernetesSavePod struct {
	Pod
	Storage
	EtcdAddress string
}

type BacupDevopsSavePod struct {
	BasePod Pod

	BaseStorage  Storage
	MinioAddress string
	MinioBucket  string
	MongoAddress string
	MysqlAddress string

	MinioFile string
	MongoFile string
	MysqlFile string
}

type RestoreDevopsPod struct {
	Pod
}

//CreatePod 创建 BacupKubernetesSavePod
func (b *BacupKubernetesSavePod) CreatePod() {
	newArgs := []string{"cluster-etcd-endpoint=" + b.EtcdAddress}

	b.Args = append(b.Args, newArgs...)
	fmt.Println("创建 BacupKubernetesSavePod 结构体", b.Args)
}

func (b *BacupKubernetesSavePod) StoragePattern(storeType int) {
	switch storeType {
	case SaveToMinio:
		newArgs := []string{"store-minio-endpoint=" + b.MinioEndpoint, "store-minio-bucket=" + b.MinioBucket, "store-minio-file=" + b.MinioFile}
		b.Args = append(b.Args, newArgs...)

		//因为是保存到minio，所以需要挂载volume
		b.VolumeMounts = map[string]string{"/backup/kubernetes/store/minio/": "minio-secret"}
		b.Volumes = map[string]string{"minio-secret": "store-minio-creds-config"}
	case SaveToSftp:
		newArgs := []string{"store-sftp-endpoint=" + b.SftpEndpoint, "store-sftp-file=" + b.SftpFilePath}
		b.Args = append(b.Args, newArgs...)

		b.VolumeMounts = map[string]string{"/backup/kubernetes/store/sftp/": "sftp-secret"}
		b.Volumes = map[string]string{"sftp-secret": "store-sftp-creds-config"}

	case SaveToLocal:
		newArgs := []string{"store-volume-file=" + b.VolueFile}
		b.Args = append(b.Args, newArgs...)
		b.IsHostPath = true

		b.VolumeMounts = map[string]string{"/newbackup": "store-out-data"}
		b.Volumes = map[string]string{"/etc/backup/data": "store-out-data"}
	}

	fmt.Println("设置 BacupKubernetesSavePod的存储方式", b.Args)
}

//CreatePod 创建 BacupDevopsSavePod
func (b *BacupDevopsSavePod) CreatePod() {
	newArgs := []string{"cluster-minio-endpoint=", b.MinioAddress, "cluster-minio-bucket=" + b.MinioBucket, "cluster-mongo-endpoint=", b.MongoAddress, "cluster-mysql-endpoint=", b.MysqlAddress}

	b.BasePod.Args = append(b.BasePod.Args, newArgs...)
	fmt.Println("创建 BacupDevopsSavePod 结构体", b.BasePod.Args)
}

func (b *BacupDevopsSavePod) StoragePattern(storeType int) {

}

func main() {
	var podTemplate PodTemplate

	volume := map[string]string{"/backup/kubernetes/cluster/etcd/": "cluster-etcd-secret"}

	podTemplate = &BacupKubernetesSavePod{
		Pod: Pod{AppName: "backup", Namespace: "kube-system", ImageName: "registry.cn-hangzhou.aliyuncs.com/kubesphere/etcd-backup:latest"},
		//AppName:     "backup",
		//ImageName:   "registry.cn-hangzhou.aliyuncs.com/kubesphere/etcd-backup:latest",
		EtcdAddress: "https://127.0.0.1:2379",
		Storage:     Storage{VolueFile: "/newbackup/jobone.zip", VolumeMounts: volume},
	}
	podTemplate.CreatePod()
	podTemplate.StoragePattern(SaveToLocal)

	tmpl, err := template.ParseFiles("./CreateBackupTool.yaml")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	var bs bytes.Buffer
	tmpl.Execute(&bs, podTemplate)

	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	// Decode the YAML file into an unstructured object
	obj := &corev1.Pod{}
	_, _, err = decoder.Decode(bs.Bytes(), nil, obj)

	if err != nil {
		panic(err)

	}

	klog.Infof(string(bs.Bytes()))

	//var devopsPodTemplate PodTemplate
	//devopsPodTemplate = &BacupDevopsSavePod{MinioAddress: "127.0.0.1:9001", MinioBucket: "bucket", MysqlAddress: "127.0.0.1:3306", MongoAddress: "127.0.0.1:2706"}
	//devopsPodTemplate.CreatePod()
	//devopsPodTemplate.StoragePattern(SaveToMinio)
}
