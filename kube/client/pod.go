package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreatePod 创建pod
func CreatePod() {
	//pod模版
	newPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-nginx",
			Labels: map[string]string{
				"yunhang-platform/backup": "backup",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Name: "my-nginx", Image: "nginx:latest", Command: []string{"sleep", "1000"}},
			},
		},
	}

	//创建pod
	pod, err := kubeClientSet.CoreV1().Pods("yunhang").Create(context.Background(), newPod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created pod %q.\n", pod.GetObjectMeta().GetName())
}
