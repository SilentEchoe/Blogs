package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	logOptions := &v1.PodLogOptions{
		Follow: false, // 对应kubectl logs -f参数
	}

	request := clientset.CoreV1().Pods("default").GetLogs("space-64d47c944c-2xfbs", logOptions)
	readCloser, err := request.Stream(context.TODO())

	r := bufio.NewReader(readCloser)
	for {
		bytes, err := r.ReadBytes('\n')
		fmt.Println(string(bytes))
		if err != nil {
			if err != io.EOF {
				return
			}
			return
		}
	}
	return
}
