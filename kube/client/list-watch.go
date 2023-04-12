// Import necessary packages
package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/kai/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	// Create a new clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Define namespace and pod name
	namespace := "kube-system"
	podName := "my-nginx"

	// Create a new context
	ctx := context.Background()

	// List options for pod
	listOptions := metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", podName),
	}

	// Watch for changes to the pod
	podWatch, err := clientset.CoreV1().Pods(namespace).Watch(ctx, listOptions)
	if err != nil {
		panic(err.Error())
	}

	// Loop through events
	for event := range podWatch.ResultChan() {
		// Check event type
		switch event.Type {
		case watch.Added:
			fmt.Printf("Pod %s added\n", podName)
		case watch.Modified:
			fmt.Printf("Pod %s modified\n", podName)
		case watch.Deleted:
			fmt.Printf("Pod %s deleted\n", podName)
		case watch.Error:
			// Check if error is a not found error
			if event.Object == nil {
				fmt.Printf("Pod %s not found\n", podName)
			} else {
				// Print error message
				fmt.Printf("Error: %v\n", event.Object)
			}
		}
	}
}
