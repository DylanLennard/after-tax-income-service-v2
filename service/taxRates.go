package service

// This file exists in the mean time while we do not have
// full state data or any sort of formal data store
// next step may be to create a JSON file with all of this info in it
// and access it dynamically via another parameter
// but for now, this will do

import "github.com/DylanLennard/after-tax-income-service-v2/model"

// StateTaxes is temporary data used since we don't have full
// state tax data
var StateTaxes model.Taxes = model.Taxes{
	Name: "State",
	Rates: []float64{0.01, 0.02, 0.04, 0.06, 0.08, 0.093,
		0.103, 0.113, 0.123, 0.133},
	Brackets: []float64{0, 8015, 19001, 29989, 41629, 52612, 268750,
		322499, 537498, 1000000},
}

// FederalTaxes is the actual federal tax data from 2018, but for now I don' t
// have the data structured in a clean way so this will have to do
var FederalTaxes model.Taxes = model.Taxes{
	Name:     "Federal",
	Rates:    []float64{0.1, 0.12, 0.22, 0.24, 0.32, 0.35, 0.37},
	Brackets: []float64{0, 9526, 38701, 82501, 157501, 200001, 500001},
}
