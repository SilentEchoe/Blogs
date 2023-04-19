package main

import (
	"bytes"
	"flag"
	"fmt"
	"path/filepath"
	"text/template"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/klog/v2"
)

type BackuoPod struct {
	AppName          string
	Namespace        string
	ImageName        string
	Args             []string
	InMongoEndpoint  string
	InMysqlEndpoint  string
	InMinioEndpoint  string
	InMinioBucket    string
	OutMinioEndpoint string
	OutMinioBucket   string
	OutMinioFile     string
	OutSftpEndpoint  string
	OutSftpFile      string
	OutVolumeFile    string
	MasterIps        []string
}

var kubeClientSet *kubernetes.Clientset

func Init() {
	kubeClientSet = newKubeClientSet()
	if kubeClientSet == nil {
		panic("init kube client failed")
	}
}

func main() {
	tmpl, err := template.ParseFiles("./backup.yaml")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	pod := BackuoPod{
		AppName:         "backup-adug16516",
		Namespace:       "kube-system",
		ImageName:       "nginx:latest",
		InMinioBucket:   "bucket",
		InMongoEndpoint: "mongodb://root:zadig@kr-mongodb:27017",
		MasterIps:       []string{"devops", "save", "in-mongo-endpoint=mongo://127.0.0.1:2710"},
	}

	var bs bytes.Buffer
	tmpl.Execute(&bs, pod)

	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	// Decode the YAML file into an unstructured object
	obj := &corev1.Pod{}
	_, _, err = decoder.Decode(bs.Bytes(), nil, obj)
	if err != nil {
		panic(err)

	}
	//klog.Infof("%#v", obj)

	klog.Infof(string(bs.Bytes()))

}

func newKubeClientSet() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 使用kubeconfig中的当前上下文,加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset

}
