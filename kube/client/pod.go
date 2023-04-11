package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
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

	//pod模版
	newPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-nginx",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Name: "my-nginx", Image: "nginx:latest", Command: []string{"sleep", "1000"}},
			},
		},
	}

	//创建pod
	pod, err := clientset.CoreV1().Pods("kube-system").Create(context.Background(), newPod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created pod %q.\n", pod.GetObjectMeta().GetName())
}

func watch() {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/path/to/kubeconfig")
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// create the informer factory
	informerFactory := informers.NewSharedInformerFactory(clientset, 0)

	// create the channel for receiving notifications
	stopper := make(chan struct{})
	defer close(stopper)

	// create the informer for pods
	podInformer := informerFactory.Core().V1().Pods().Informer()

	// add event handlers for when pods are added, updated, or deleted
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("Pod added: %s\n", obj.(*corev1.Pod).GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("Pod updated: %s\n", newObj.(*corev1.Pod).GetName())
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("Pod deleted: %s\n", obj.(*corev1.Pod).GetName())
		},
	})

	// start the informer
	go podInformer.Run(stopper)

	// wait for the channel to be closed
	<-stopper
}
