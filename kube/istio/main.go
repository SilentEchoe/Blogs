package main

import (
	"flag"
	"fmt"
	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	istioClient := NewIstioClient()
	fmt.Println("连接成功:", istioClient)
}

func NewIstioClient() *istio.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// 使用kubeconfig中的当前上下文,加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	//创建Istio连接
	istioClient, err := istio.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//todo

	return istioClient

}
