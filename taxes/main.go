package main

import (
	"example.com/taxes/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		taxIncludedPriceJob := prices.NewTaxIncludedPriceJob(taxRate)
		taxIncludedPriceJob.Process()
	}

}
