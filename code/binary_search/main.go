package main

func main() {

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var key [26]int
		for _, c := range str {
			key[c-'a']++
		}
		m[key] = append(m[key], str)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
