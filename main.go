package main

type Test struct {
	value int
}

func main() {
	var slice1 = make([]Test, 0)
	test := Test{
		value: 1,
	}
	slice1 = append(slice1, test)

	for key, value := range slice1 {
		println(key)
		println(value.value)
	}

}
