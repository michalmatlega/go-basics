package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moreNumbers := []int{5, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&moreNumbers)
	transformedNumbers := transformNumbers(&moreNumbers, transformerFn1)
	fmt.Println(transformedNumbers)

	transformerFn2 := getTransformerFunction(&numbers)
	transformedNumbers2 := transformNumbers(&numbers, transformerFn2)
	fmt.Println(transformedNumbers2)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}

	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
