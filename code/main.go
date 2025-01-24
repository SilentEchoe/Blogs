package main

import (
	"fmt"
	"sync"
	"time"
)

type Rule struct {
	Endpoint            string   `yaml:"target-endpoint"`
	AcquisitionInterval int      `yaml:"acquisition-interval"`
	MaxWorkSum          int      `yaml:"max-work-sum"`
	Timeout             int      `yaml:"timeout-seconds"`
	Actions             []Action `yaml:"actions"`
}

type Action struct {
	Agreement string `yaml:"agreement"`
	Value     string `yaml:"value"`
}

func main() {
	//data, err := ioutil.ReadFile("/Users/kai/PrivateProject/src/Blogs/code/test.yaml")
	//if err != nil {
	//	log.Fatalf("无法读取 YAML 文件: %v", err)
	//}
	//
	//var rules []Rule
	//err = yaml.Unmarshal(data, &rules)
	//if err != nil {
	//	log.Fatalf("无法解析 YAML 文件: %v", err)
	//}
	//
	//fmt.Println(rules)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("test", i)
			time.Sleep(2 * time.Second)

		}()
	}

	wg.Wait()

}
