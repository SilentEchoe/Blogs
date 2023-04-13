package main

import (
	"bytes"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/klog/v2"
	"text/template"
)

type BackuoPod struct {
	AppName          string
	Namespace        string
	ImageName        string
	Args             []string
	InMongoEndpoint  string
	InMysqlEndpoint  string
	InMinioEndpoint  string
	InMinioBucket    string
	OutMinioEndpoint string
	OutMinioBucket   string
	OutMinioFile     string
	OutSftpEndpoint  string
	OutSftpFile      string
	OutVolumeFile    string
}

func main() {
	tmpl, err := template.ParseFiles("./backup-devops.yaml")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	pod := BackuoPod{
		AppName:         "backup-adug16516",
		Namespace:       "kube-system",
		ImageName:       "nginx:latest",
		InMinioBucket:   "bucket",
		InMongoEndpoint: "mongodb://root:zadig@kr-mongodb:27017",
	}

	var bs bytes.Buffer
	tmpl.Execute(&bs, pod)

	decoder := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	// Decode the YAML file into an unstructured object
	obj := &corev1.Pod{}
	_, _, err = decoder.Decode(bs.Bytes(), nil, obj)
	if err != nil {
		panic(err)
	}
	klog.Infof("%#v", obj)

}
