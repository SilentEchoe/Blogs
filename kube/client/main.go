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

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

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
