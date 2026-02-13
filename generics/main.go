package main

import "fmt"

func main() {
	result := add(1, 2.1)
	fmt.Println(result)

	result2 := add(1.1, 2.2)
	fmt.Println(result2)

	result3 := add("foo", "bar")
	fmt.Println(result3)

	//result4 := add[bool](true, false)
	//fmt.Println(result4)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

//func add[T bool](a, b T) T {
//	return a + b
//}
