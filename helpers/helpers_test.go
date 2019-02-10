package helpers

import "testing"

func TestCalculateOtherTaxes(t *testing.T){

	// setup vars
	var income float64 = 50000
	selfEmpStatus := true

	// expected vals
	expectedOtherTaxesSelfEmp := 7650.00
	expectedOtherTaxesNotSelfEmp := 3825.00

	// test withs self employment status is true first
	otherTaxes := CalculateOtherTaxes(income, selfEmpStatus)
	if otherTaxes != expectedOtherTaxesSelfEmp {
		t.Error("Expected", expectedOtherTaxesSelfEmp, "got", otherTaxes)
	}

	// then test with self employment status is false
	otherTaxes = CalculateOtherTaxes(income, !selfEmpStatus)
	if otherTaxes != expectedOtherTaxesNotSelfEmp {
		t.Error("Expected ", expectedOtherTaxesNotSelfEmp, "got", otherTaxes)
	}
}