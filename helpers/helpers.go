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

// GetTaxInfo gets all the full data
func GetTaxInfo() (model.Taxes, model.Taxes) {

	// open the two files and unmarshal the JSON appropriately
	fedTaxFile, errFed := ioutil.ReadFile("data/FederalTaxData.json")
	if errFed != nil {
		fmt.Println(errFed)
	}
	stateTaxFile, errState := ioutil.ReadFile("data/StateTaxData.json")
	if errState != nil {
		fmt.Println(errState)
	}

	//unmarshal the JSON
	var FederalTaxes model.Taxes
	var StateTaxes model.Taxes
	errUnMarshFed := json.Unmarshal(fedTaxFile, &FederalTaxes)
	if errUnMarshFed != nil {
		fmt.Println(errUnMarshFed)
	}
	errUnMarshState := json.Unmarshal(stateTaxFile, &StateTaxes)
	if errUnMarshState != nil {
		fmt.Println(errUnMarshState)
	}

	fmt.Println(StateTaxes)

	return FederalTaxes, StateTaxes

}
