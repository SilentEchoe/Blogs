package main

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type DepHandler struct {
}

func (d *DepHandler) OnAdd(obj interface{}) {}
func (d *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	if dep, ok := newObj.(*v1.Deployment); ok {
		fmt.Println(dep.Name)
	}
}
func (d *DepHandler) OnDelete(obj interface{}) {
}
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
		panic(err.Error())
	}

	// c Getter, resource string, namespace string, fieldSelector fields.Selector
	var list_watch = cache.NewListWatchFromClient(restClient, "deployments", "kube-system", fields.Everything())
	s, c := cache.NewInformer(
		list_watch,
		&v1.Deployment{},
		0,
		&DepHandler{},
	)
	c.Run(wait.NeverStop)
	s.List()
}
