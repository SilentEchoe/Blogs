package main

func main() {
	//使用Map 表示一个映射的关系表
	graph := map[string][]string{}

	// 构建出'我'的关系图网
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}

}

func search(name string) {
	//使用一个数组的队列来存放
	var searchQueue []string
	searchQueue = append(searchQueue, name)
	var searched []string // 用于记录检查过的人

	for _, v := range searchQueue {
		if !isexist(searched, v) {

		}

	}

}

// 是否存在
func isexist(searched []string, name string) bool {
	for _, v := range searched {
		if v == name {
			return true
		}
	}
	return false
}
