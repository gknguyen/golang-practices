package prices

import (
	"fmt"
	"priceCalculator/conversion"
	"priceCalculator/filemanager"
)

type TaxPriceJob struct {
	TaxRate     float64
	InputPrices []float64
	TaxPrices   map[string]string
	IOManager   filemanager.FileManager
}

func NewTaxPriceJob(taxRate float64, fm filemanager.FileManager) *TaxPriceJob {
	return &TaxPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}

func (job *TaxPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxPrices = result
	err := job.IOManager.WriteJSON(job)
	if err != nil {
		fmt.Println(err)
	}
}

func (job *TaxPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}
