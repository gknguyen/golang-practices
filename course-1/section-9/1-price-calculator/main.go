package main

import (
	"fmt"
	"priceCalculator/filemanager"
	"priceCalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("abc.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()

		taxRateJob := prices.NewTaxPriceJob(taxRate, fm)
		err := taxRateJob.Process()
		if err != nil {
			fmt.Println(err)
		}
	}

}
