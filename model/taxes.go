package model

import "math"

type Taxes struct {
	Brackets []float64
	Rates []float64
	Name string
}


// Calculates your after tax income for a given set of taxes 
func (T Taxes) Calculate(income float64) float64{

	var amount float64

	for i, val := range T.Brackets{
		if i == 0 {
			continue
		} else if (i == (len(T.Brackets)) - 1) && (income > val){
			amount += (income-val) * T.Rates[i]
		} else if income > val {
			amount += (val - T.Brackets[i-1]) * T.Rates[i-1]
		} else {
			amount += (income - T.Brackets[i-1]) * T.Rates[i-1]
			break
		}
	}
	// round to nearest two cents
	amount = math.Round(amount*100)/100
	return amount
}