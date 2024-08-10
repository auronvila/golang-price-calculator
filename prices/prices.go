package prices

import (
	"fmt"
	typeConverter "github.com/price-calculator/converter"
	"github.com/price-calculator/fileManager"
	"log"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) loadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An err occurred, Could not open the file.")
		fmt.Println(err)
		file.Close()
		return
	}

	lines, err := typeConverter.FileReaderAndParser(file)
	if err != nil {
		fmt.Println("Error in parsing data from file", err)
		file.Close()
	}

	job.InputPrices = lines
	file.Close()
}

func (job TaxIncludedPriceJob) Process() {
	job.loadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	outPutFileName := fmt.Sprintf("result_%.0f.json", job.TaxRate*100)
	managedFile := fileManager.New(outPutFileName)
	err := managedFile.WriteJson(job)
	if err != nil {
		log.Fatal(err)
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
