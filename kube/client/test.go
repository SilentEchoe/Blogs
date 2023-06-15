package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {

	client := newKube()

	// 获取当前 Pod 的信息
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}
	pod, err := client.CoreV1().Pods(os.Getenv("POD_NAMESPACE")).Get(
		context.TODO(),
		hostname,
		metav1.GetOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// 解析当前 Pod 所属的 Service 名称
	service := ""
	for _, env := range pod.Spec.Containers[0].Env {
		if env.Name == "MY_POD_SERVICE_NAME" {
			service = env.Value
			break
		}
	}
	if service == "" {
		log.Fatal("MY_POD_SERVICE_NAME environment variable not found")
	}

	// 查询对应的 Service
	svc, err := client.CoreV1().Services(os.Getenv("POD_NAMESPACE")).Get(
		context.TODO(),
		service,
		metav1.GetOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// 输出 Service 的信息
	fmt.Printf("Service: %s/%s\n", svc.Namespace, svc.Name)
	for _, port := range svc.Spec.Ports {
		fmt.Printf("\tPort: %d (targetPort: %d)\n", port.Port, port.TargetPort.IntVal)
	}
}

func newKube() *kubernetes.Clientset {
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

	// 创建clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset

}
