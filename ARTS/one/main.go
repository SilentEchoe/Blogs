/* ARTS 活动，每周写一个算法题，读一篇英文文章，分享一个小技术，分享一个观点 */
package main

//本周算法题是数组相关
//增减字符串匹配
//由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:
//如果 perm[i] < perm[i + 1] ，那么 s[i] == 'I'
//如果 perm[i] > perm[i + 1] ，那么 s[i] == 'D'
//给定一个字符串 s ，重构排列 perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。
//输入：s = "IDID"
//输出：[0,4,1,3,2]
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/di-string-match
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	diStringMatch("IDID")
}

func diStringMatch(s string) []int {
	var parm []int

	return parm
}
