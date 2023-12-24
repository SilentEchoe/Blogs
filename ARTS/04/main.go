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
