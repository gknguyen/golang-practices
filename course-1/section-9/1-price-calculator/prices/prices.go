package prices

import (
	"fmt"
	"priceCalculator/conversion"
	"priceCalculator/iomanager"
)

type TaxPriceJob struct {
	TaxRate     float64             `json:"tax_rate"`
	InputPrices []float64           `json:"input_prices"`
	TaxPrices   map[string]string   `json:"taxed_prices"`
	IOManager   iomanager.IOManager `json:"-"`
}

func NewTaxPriceJob(taxRate float64, iom iomanager.IOManager) *TaxPriceJob {
	return &TaxPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

func (job *TaxPriceJob) Process(doneChan chan bool, errorChan chan error) error {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxPrices = result
	err = job.IOManager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return err
	}

	doneChan <- true
	return nil
}

func (job *TaxPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}
