package main

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {

	endpoint := "localhost:9090"
	accessKeyID := "AKIAIOSFODNN7EXAMPLE"
	secretAccessKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln("创建 MinIO 客户端失败", err)
		return
	}
	log.Printf("创建 MinIO 客户端成功")

	if err != nil {
		fmt.Println("错误信息", err)
	}

	listBuckets, _ := minioClient.ListBuckets()
	fmt.Println(listBuckets)

	// 创建一个叫 mybucket 的存储桶。
	bucketName := "mybucket"
	location := ""

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", bucketName)
		} else {
			log.Fatalln("查询存储桶状态异常", err, exists)
		}
	}
	log.Printf("创建存储桶 %s 成功", bucketName)
}
