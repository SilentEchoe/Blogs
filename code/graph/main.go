package main

import "fmt"

/*
	图算法相关

图模拟一组连接，由节点和边组成。一个节点可能与众多节点直接相连。
*/
type GraphMap map[string][]string

func main() {
	var graphMap GraphMap = make(GraphMap, 0)
	graphMap["you"] = []string{"alice", "bob", "claire"}
	graphMap["bob"] = []string{"anuj", "peggy"}
	graphMap["alice"] = []string{"peggy"}
	graphMap["claire"] = []string{"tom", "johnny"}
	graphMap["anuj"] = []string{}
	graphMap["peggy"] = []string{}
	graphMap["tom"] = []string{}
	graphMap["johnny"] = []string{}

	search_queue := graphMap["you"]

	for {
		if len(search_queue) > 0 {
			var person string
			person, search_queue = search_queue[0], search_queue[1:]
			if personIsTom(person) {
				fmt.Printf("%s is already in the queue for you.\n", person)
				break
			} else {
				search_queue = append(search_queue, graphMap[person]...)
			}
		} else {
			fmt.Println("Not found in search queue")
			break
		}
	}
}

func personIsTom(p string) bool {
	return p == "tom"
}
