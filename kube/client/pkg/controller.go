package pkg

import (
	informer "k8s.io/client-go/informers/core/v1"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	lister "k8s.io/client-go/listers/core/v1"
	netlister "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
)

type controller struct {
	client        kubernetes.Interface
	ingressLister netlister.IngressLister
	serviceLister lister.ServiceLister
}

func (c *controller) Run(stopCn chan struct{}) {
	<-stopCn
}

func (c *controller) addService(obj interface{}) {

}

func (c *controller) updateService(oldobj interface{}, newobj interface{}) {

}

func (c *controller) deleteIngress(obj interface{}) {

}

func NewController(client kubernetes.Interface, serviceInformer informer.ServiceInformer, ingressInformer netInformer.IngressInformer) controller {
	c := controller{
		client:        client,
		serviceLister: serviceInformer.Lister(),
		ingressLister: ingressInformer.Lister(),
	}

	// 注册方法
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})
	return c
}
