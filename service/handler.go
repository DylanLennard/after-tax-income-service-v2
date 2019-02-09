package service 

import (
	"fmt"
	"math"
	"github.com/DylanLennard/after-tax-service-v2/helpers"
)

// This handler results in the calculation of your after tax income
// assuming you live in the state of california for the
// 2018 tax year. 
// It'll eventually take in the URL and parse out your income and status level 
// and pass that into the Tax function defined in helpers  
// AFter that it calculates federal, state, and other taxes
// and subtracts them from your income and returns the final income
func AfterTaxIncome(income float64, status bool) float64 {
	fmt.Println("Looking good, bruh")
	federalTax := FederalTaxes.Calculate(income)
	stateTax := StateTaxes.Calculate(income)
	otherTax := helpers.CalculateOtherTaxes(income, status)
	afterTax := income - (federalTax + stateTax + otherTax)
	return math.Round(afterTax*100)/100
}
