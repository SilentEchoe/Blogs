package main

import (
	"fmt"
	"net/http"

	"LearningNotes-Go/routers"

	"LearningNotes-Go/pkg/setting"
)

func main() {

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
