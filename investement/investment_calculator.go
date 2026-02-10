package main

import (
	"fmt"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value (adjusted for Inflation): %.1f", futureRealValue)

	//fmt.Printf(`Future Value: %.0f\n
	//
	//		Future real value: %.0f`, futureValue, futureRealValue)
	//fmt.Println("Future value:", futureValue)
	//fmt.Printf("Future Value: %.0f\nFuture real value: %.0f", futureValue, futureRealValue)
	//fmt.Println("Future real value:", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

}

func outputText(text, text2 string) {
	fmt.Print(text)
	fmt.Print(text2)
}
