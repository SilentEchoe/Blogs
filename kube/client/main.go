package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	client()
}

//func restClient() {
//	// config 可以从指定的文件中获取，例如：clientcmd.RecommendedHomeFile 从 home 文件中
//	// 如果不传递参数，它会从classname 中获取
//	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
//	if err != nil {
//		fmt.Println(err)
//		panic(err)
//	}
//
//	config.GroupVersion = &v1.SchemeGroupVersion
//	config.NegotiatedSerializer = scheme.Codecs
//	config.APIPath = "/api"
//
//	// client
//	restClient, err := rest.RESTClientFor(config)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// get data
//	pod := v1.Pod{}
//	err = restClient.Get().Namespace("ingress-nginx").Resource("pods").Name("ingress-nginx-admission-create-w8sd2").Do(context.Background()).Into(&pod)
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Println(pod.Name)
//	}
//}

func client() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	coreV1 := clientset.CoreV1()
	pod, err := coreV1.Pods("default").Get(context.Background(), "test", v1.GetOptions{})
	if err != nil {
		println(err)
	} else {
		println(pod.Name)
	}
}
