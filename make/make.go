package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) //preallocates two empty elements, performance optimizations

	userNames[0] = "Julie"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["node"] = 4.7 //now it needs to reallocate memory

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for index, value := range courseRatings {
		fmt.Println(index, value)
	}
}
