package src

import (
	"fmt"
	example_firts "github.com/AnAnonymousFriend/LearningNotes-Go/src/first"
)

func main()  {
	fmt.Println("hello")
}

func NewPersonMessage()  {
	pm := example_firts.PersonMessage{
		Id : 1234,
		IsAdult: true,
		Name: "Dave",
		LuckNumvers: []int32{1,2,3,4,5},
	}

	fmt.Println(pm)
	pm.Name = "Nick"

	fmt.Println(pm)


}