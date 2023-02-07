package main

import (
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type PodHandler struct {
}

func (p *PodHandler) OnAdd(obj interface{}) {}

func (p *PodHandler) OnDelete(obj interface{}) {}

func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	if pods, ok := newObj.(*corev1.Pod); ok {
		fmt.Println(pods.Name)
	}
}

// 工厂模式
func main() {
	var client = NewClient()
	factory := informers.NewSharedInformerFactory(client, 0)
	podinformer := factory.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(&PodHandler{})
	factory.Start(wait.NeverStop)
	select {}
}

func NewClient() *kubernetes.Clientset {
	var kubeconfig *string
	//如果是windows，那么会读取C:\Users\xxx\.kube\config 下面的配置文件
	//如果是linux，那么会读取~/.kube/config下面的配置文件
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}
