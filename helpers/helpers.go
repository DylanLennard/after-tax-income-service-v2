package helpers

import (
	"github.com/DylanLennard/after-tax-income-service-v2/model"
)

// CalculateOtherTaxes calculates SSI and medicare related taxes
func CalculateOtherTaxes(income float64, selfEmpStatus bool) float64 {

	medicareRate := 0.0145
	medicareUpperRate := 0.0235 // for those making over $200000
	ssiRate := 0.062            //up to $118,500
	var totalOtherTax, ssiTax, medicareTax float64

	if selfEmpStatus {
		medicareRate = medicareRate * 2
		medicareUpperRate = medicareUpperRate * 2
		ssiRate = ssiRate * 2
	}

	if income < 118500 {
		ssiTax = income * ssiRate
		medicareTax = income * medicareRate
	} else {
		ssiTax = 118500 * ssiRate
		if income <= 200000 {
			medicareTax = income * medicareRate
		} else {
			medicareTax = income * medicareUpperRate
		}
	}
	totalOtherTax = ssiTax + medicareTax
	return totalOtherTax
}

// GetTaxInfo gets all the full data
func GetTaxInfo() (model.Taxes, model.Taxes) {

	// open the two files and unmarshal the JSON appropriately
	// fedTaxFile, errFed := os.Open("users.json")
	// defer fedTaxFile.close()
	// if errFed != nil {
	// 	fmt.Println(errFed)
	// }
	// stateTaxFile, errState := os.Open("users.json")
	// defer stateTaxFile.close()
	// if errState != nil {
	// 	fmt.Println(errState)
	// }

	//unmarshal the JSON

	// TODO: delete this once you've gotten data from file explicitly
	StateTaxes := model.Taxes{
		Rates: []float64{0.01, 0.02, 0.04, 0.06, 0.08, 0.093,
			0.103, 0.113, 0.123, 0.133},
		Brackets: []float64{0, 8015, 19001, 29989, 41629, 52612, 268750,
			322499, 537498, 1000000},
	}
	FederalTaxes := model.Taxes{
		Rates:    []float64{0.1, 0.12, 0.22, 0.24, 0.32, 0.35, 0.37},
		Brackets: []float64{0, 9526, 38701, 82501, 157501, 200001, 500001},
	}

	return FederalTaxes, StateTaxes

}
