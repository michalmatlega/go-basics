package main

import "fmt"

type str string

// we can extend builtin string type with extra methods by using our str alias
func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Max"

	name.log()
}
