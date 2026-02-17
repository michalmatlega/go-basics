package main

import (
	"fmt"

	"example.com/taxes/filemanager"
	"example.com/taxes/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("pricesssss.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		taxIncludedPriceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := taxIncludedPriceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
