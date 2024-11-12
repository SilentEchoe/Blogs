package main

import "fmt"

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}

// 找出由相同字母组成的单词
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	oldMap := NewAnagramsMap()
	newMap := make(map[int][]string)
	rest := make([][]string, 0)
	for _, v := range strs {
		primeNumber := 1
		for _, j := range v {
			primeNumber *= oldMap[string(j)]
		}
		// 这样就得到一个单词的质数
		newMap[primeNumber] = append(newMap[primeNumber], v)
	}

	for _, v := range newMap {
		//fmt.Println(v)
		rest = append(rest, v)
	}

	return rest
}

// NewAnagramsMap 用质数表示26个字母，把字符串的各个字母相乘以，字母异位词的乘积必定是相等的
func NewAnagramsMap() map[string]int {
	m := map[string]int{
		"a": 2,
		"b": 3,
		"c": 5,
		"d": 7,
		"e": 11,
		"f": 13,
		"g": 17,
		"h": 19,
		"i": 23,
		"j": 29,
		"k": 31,
		"l": 37,
		"m": 41,
		"n": 43,
		"o": 47,
		"p": 53,
		"q": 59,
		"r": 61,
		"s": 67,
		"t": 71,
		"u": 73,
		"v": 79,
		"w": 83,
		"x": 89,
		"y": 97,
		"z": 101,
	}

	return m
}
