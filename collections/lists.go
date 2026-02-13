package main

import (
	"fmt"
	"slices"
)

func main() {
	biz()
	bar()
}

func biz() {
	hobbies := [3]string{"Hiking", "Biking", "Gamedev"}
	fmt.Println("1", hobbies)

	fmt.Println("2a", hobbies[0])
	fmt.Println("2b", [2]string{hobbies[1], hobbies[2]})

	fmt.Println("3a", hobbies[0:2])
	firstAndSecondHobbies := hobbies[:2]
	fmt.Println("3b", firstAndSecondHobbies)
	fmt.Println("4", firstAndSecondHobbies[1:3])

	courseGoals := []string{"Learn GO", "Find job with GO"}
	courseGoals[1] = "Find fantastic job"
	courseGoals = append(courseGoals, "Get better salary")
	fmt.Println(courseGoals)

	type Product struct {
		id    int
		title string
		price float64
	}

	products := []Product{
		{id: 1, title: "Apple", price: 12.99},
		{id: 2, title: "Banana", price: 13.99},
	}

	newProducts := slices.Clone(append(products, Product{id: 1, title: "Oregano", price: 10.99}))
	fmt.Println(newProducts)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		a) The first element (standalone)
//		b_ The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		a) the first and second elements.
//		b) Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func bar() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	//prices[2] = 11.99	//cant

	prices = append(prices, 5.99)
	fmt.Println(prices)

	slicePrice := prices[1:]
	appendedSlice := slices.Clone(append(slicePrice, 21.37))
	appendedSlice[0] = 0.99
	fmt.Println(appendedSlice, prices)

	fmt.Println("unpacking", append(prices, []float64{4.5, 6.7, 8.9}...))
}

func foo() {
	var productNames = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:3] //slice
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3] //slice
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:] //slice
	featuredPrices[0] = 199.9   //slice wplywa na oryginalna tablice!
	highlightedPrices := featuredPrices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) //you can reslice only to the right!! you can select more from a smaller slice
}
