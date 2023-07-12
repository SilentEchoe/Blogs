package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var kubeClientSet *kubernetes.Clientset

func init() {
	kubeClientSet = newKubeClientSet()
	if kubeClientSet == nil {
		panic("init kube client failed")
	}
}

func main() {
	svc := GenerateAgentService()
	newSvc, err := kubeClientSet.CoreV1().Services("default").Create(context.Background(), svc, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newSvc.Spec.Ports[0].NodePort)
}

func GenerateAgentService() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-service",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "agent",
			},
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.ProtocolTCP,
					Port:     80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 8000,
					},
				},
			},
		},
	}
}

func newKubeClientSet() *kubernetes.Clientset {
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
