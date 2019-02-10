package service


import "github.com/DylanLennard/after-tax-income-service-v2/model"

var StateTaxes model.Taxes = model.Taxes{
	Name: "State", 
	Rates: []float64{0.01, 0.02, 0.04, 0.06, 0.08, 0.093,
		             0.103, 0.113, 0.123, 0.133},
	Brackets: []float64{0, 8015, 19001, 29989, 41629, 52612, 268750,
					    322499, 537498, 1000000},
}


var FederalTaxes model.Taxes = model.Taxes{
	Name: "Federal", 
	Rates: []float64{0.1, 0.12, 0.22, 0.24, 0.32, 0.35, 0.37},
	Brackets: []float64{0, 9526, 38701, 82501, 157501, 200001, 500001},
}
