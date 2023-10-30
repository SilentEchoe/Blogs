package main

import (
	"flag"
	"fmt"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned"
	verleroInformer "github.com/vmware-tanzu/velero/pkg/generated/informers/externalversions"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
	"time"
)

var clientSetOnce = sync.Once{}
var clientSet versioned.Interface

func main() {

	GetVeleroClient()

	// 初始化 informer factory (一个小时List 一次)
	informerFactory := verleroInformer.NewSharedInformerFactory(clientSet, time.Hour)

	// 对Backup资源进行监听
	backupInformer := informerFactory.Velero().V1().Backups()

	// 创建 Informer（相当于注册到工厂中去，这样下面启动的时候就会去 List & Watch 对应的资源）

	informer := backupInformer.Informer()

	// 创建 Lister
	backupLister := backupInformer.Lister()

	// 注册时间处理程序
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    OnAdd,
		UpdateFunc: OnUpdate,
		DeleteFunc: OnDelete,
	})

	stopper := make(chan struct{})
	defer close(stopper)

	// 启动 informer，List & Watch
	informerFactory.Start(stopper)

	// 等待所有启动的 Informer 的缓存被同步
	informerFactory.WaitForCacheSync(stopper)

	// 从本地缓存中获取 default 中的所有 deployment 列表
	backups, err := backupLister.Backups("velero").List(labels.Everything())
	if err != nil {
		panic(err)
	}
	for idx, backup := range backups {
		fmt.Printf("%d -> %s\\n", idx+1, backup.Name)
	}
	<-stopper
}

func OnAdd(obj interface{}) {
	backup := obj.(*v1.Backup)
	fmt.Println("add a backup:", backup.Name)
}

func OnUpdate(old, new interface{}) {
	backup := new.(*v1.Backup)
	if backup.Status.Progress != nil {
		result := fmt.Sprintf("当前进度%d/%d", backup.Status.Progress.ItemsBackedUp, backup.Status.Progress.TotalItems)
		fmt.Println(result)
	}
	fmt.Println("当前状态：", backup.Status.Phase)
}

func OnDelete(obj interface{}) {

}

func GetVeleroClient() versioned.Interface {
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

	// 创建 veleroClient
	veleroClient, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	} else {
		clientSet = veleroClient
	}

	return clientSet
}
