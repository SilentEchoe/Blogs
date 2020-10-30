package main

import (
	"fmt"
	_ "strings"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || haystack == needle {
		return 0
	}

	if len(needle) > 0 && len(haystack) > 0 {
		for i := 0; i <= len(haystack)-len(needle); i++ {
			ch := haystack[i : i+len(needle)]
			if ch == needle {
				return i
			}
			fmt.Printf("%q ", ch)
		}
	}
	return -1
}

func main() {
	var a = "abc"
	var b = "c"
	println(strStr(a, b))
}
