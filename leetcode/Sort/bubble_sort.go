package main

func main() {

}

func BubbleSort(sum []int) []int {
	for i := 0; i < len(sum)-1; i++ {
		for j := 0; j < len(sum)-1-i; j++ {
			if sum[j] > sum[j+1] {
				var tmp = sum[j]
				sum[j+1] = sum[j]
				sum[j] = tmp
			}
		}
	}
	return sum
}
