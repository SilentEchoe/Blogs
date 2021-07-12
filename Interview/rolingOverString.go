package main

import "fmt"

func main() {

	fmt.Println(RolingOver("123456789"))
}

func RolingOver(s string) (string, bool) {
	str := []rune(s)
	l := len(str)

	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str), true
}
