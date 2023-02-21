package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type KafkaCluster struct {
	Rule map[string]string `yaml:"rule"`
	test map[string]string
}

func main() {
	var c KafkaCluster
	////读取yaml配置文件, 将yaml配置文件，转换struct类型
	conf := c.getConf()

	rules := []string{"/api/v1/home/app", "/api/v1/home/log"}
	test2 := getRuleKeyByPath(rules, conf)

	if _, ok := conf.Rule["get_app"]; ok {
		fmt.Println("存在ok")
	}

	fmt.Println("该用户可访问的接口为:", test2)

}

func getRuleKeyByPath(rule []string, conf *KafkaCluster) map[string]string {
	result := make(map[string]string)

	for _, val := range rule {
		for rulekey, values := range conf.Rule {
			if values == val {
				result[rulekey] = values
			}
		}
	}

	return result
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
