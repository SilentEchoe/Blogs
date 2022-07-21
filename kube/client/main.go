package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config 可以从指定的文件中获取，例如：clientcmd.RecommendedHomeFile 从 home 文件中
	// 如果不传递参数，它会从classname 中获取
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		fmt.Println(err)
	}

	// get data
	pod := v1.Pod{}
	err = restClient.Get().Namespace("ingress-nginx").Resource("pods").Name("ingress-nginx-admission-create-w8sd2").Do(context.Background()).Into(&pod)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pod.Name)
	}
}
