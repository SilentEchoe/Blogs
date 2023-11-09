package main

import (
	million "LearningNotes-Go/ARTS/01/million"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	dispatcher := million.NewDispatcher(million.MaxWorker)
	dispatcher.Run()
	// 定义路由
	r.POST("upload", upload)
	r.Run("localhost:8000")
}

func upload(c *gin.Context) {
	// TODO 解析传递过来的信息

	req := million.PayloadReq{
		WindowsVersion: "v0.1",
		Token:          "ABVCHKL",
		Payloads:       []million.Payload{{BucketName: "Demo"}},
	}

	// 解析以后将里面的Pay加入到工作池中
	for _, payload := range req.Payloads {
		work := million.Job{Payload: payload}
		million.JobQueue <- work
	}
	// 返回OK
	c.Writer.WriteHeader(http.StatusOK)
}
