package puppy

import (
	"fmt"

	"github.com/dupakarovsky/dog"
)



func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From110(){
	fmt.Println("Hello I'm a function added to version 1.1.0")
}
func From120(){
	fmt.Println("Added to version 1.2.0")
}