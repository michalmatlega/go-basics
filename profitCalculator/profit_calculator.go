package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//var revenue float64
	//var expenses float64
	//var taxRate float64

	//fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)
	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)
	//fmt.Print("TaxRate: ")
	//fmt.Scan(&taxRate)

	revenue, err := getUserInput("Revenue: ")
	check(err)

	expenses, err := getUserInput("Expenses: ")
	check(err)

	taxRate, err := getUserInput("TaxRate: ")
	check(err)

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("ebt =", earningsBeforeTax)
	fmt.Println("profit =", earningsAfterTax)
	fmt.Println("ratio =", ratio)

	results := fmt.Sprintf("%v\n%v\n%v", earningsBeforeTax, earningsAfterTax, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic("Can't get user input")
	}
}

func getUserInput(message string) (float64, error) {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)

	if userInput <= 0.0 {
		return 0.0, errors.New("you can not provide non-positive value")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, earningsAfterTax, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)
	ratio = earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
