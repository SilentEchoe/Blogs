package main

import (
	"fmt"
	"github.com/minio/minio-go/v6"
)

func main() {
	// 使用ssl
	ssl := true

	// 初使化minio client对象。
	_, err := minio.New("localhost:9000", "Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", ssl)
	if err != nil {
		fmt.Println(err)
		return
	}
}
