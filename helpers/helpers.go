package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

// CheckFileRead will return the bytestream resulting from reading a file
func CheckFileRead(fname string) []byte {
	fileStream, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
	}
	return fileStream
}

// GetTaxInfo gets all the full data
func GetTaxInfo(state string) (model.Taxes, model.Taxes) {

	// open the two files and unmarshal the JSON appropriately
	fedTaxFile := CheckFileRead("data/FederalTaxData.json")
	stateTaxFile := CheckFileRead("data/StateTaxData.json")

	// declare the vars
	var FederalTaxes model.Taxes
	var StateTaxes map[string]model.Taxes

	// unmarshal the JSON
	err := json.Unmarshal(fedTaxFile, &FederalTaxes)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(stateTaxFile, &StateTaxes)
	if err != nil {
		fmt.Println(err)
	}

	return FederalTaxes, StateTaxes[state]

}
