package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type KafkaCluster struct {
	Rule Rule `yaml:"rule"`
}

type Rule struct {
	Home map[string]string `yaml:"home"`
}

func main() {
	var c KafkaCluster
	//读取yaml配置文件, 将yaml配置文件，转换struct类型
	conf := c.getConf()

	//将对象，转换成json格式
	data, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
	//fmt.Println(data)
}

//读取Yaml配置文件,
//并转换成conf对象  struct结构
func (kafkaCluster *KafkaCluster) getConf() *KafkaCluster {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("./rule_config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	//err = yaml.Unmarshal(yamlFile, kafkaCluster)
	err = yaml.UnmarshalStrict(yamlFile, kafkaCluster)

	if err != nil {
		fmt.Println(err.Error())
	}

	return kafkaCluster
}
