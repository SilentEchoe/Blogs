package main

import (
	"fmt"
)

type Address struct {
    Province    string
    City        string
    ZipCode     int
    PhoneNumber string
}



func main()  {
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)
}

