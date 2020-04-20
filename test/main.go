package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a []int
	configFiles := "1,22,13,28,27,21,23"
	configs := strings.Split(configFiles, ",")
	for _, v := range configs {
		id, err := strconv.Atoi(v)
		if err == nil {
			a = append(a, id)
		}

	}

	fmt.Println(a)

}
