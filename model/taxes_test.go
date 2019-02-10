package model 

import "testing"

// setup vars
var StateTaxes Taxes = Taxes{
	Name: "State", 
	Rates: []float64{0.01, 0.02, 0.04, 0.06, 0.08, 0.093,
		             0.103, 0.113, 0.123, 0.133},
	Brackets: []float64{0, 8015, 19001, 29989, 41629, 52612, 268750,
					    322499, 537498, 1000000},
}

var FederalTaxes Taxes = Taxes{
	Name: "Federal", 
	Rates: []float64{0.1, 0.12, 0.22, 0.24, 0.32, 0.35, 0.37},
	Brackets: []float64{0, 9526, 38701, 82501, 157501, 200001, 500001},
}

var income float64 = 50000


// expected vals  
var expectedStateTaxes float64 = 2107.47
var expectedFedTaxes float64 = 6939.38

// tests
func TestStateTaxes(t *testing.T){
	
	stateTaxesPaid := StateTaxes.Calculate(income)
	fedTaxesPaid := FederalTaxes.Calculate(income)

	if stateTaxesPaid != expectedStateTaxes{
		t.Error("Expected", expectedStateTaxes, "got", stateTaxesPaid)
	}

	if fedTaxesPaid != expectedFedTaxes{
		t.Error("Expected", expectedFedTaxes, "got", fedTaxesPaid)
	}
}