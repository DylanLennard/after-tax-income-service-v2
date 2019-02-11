package service

import (
	"github.com/DylanLennard/after-tax-income-service-v2/helpers"
)

type Event struct {
	Income               float64 `json:"Income"`
	SelfEmploymentStatus bool    `json:"SelfEmploymentStatus"`
}

type MyResponse struct {
	AfterTaxIncome   float64 `json:"AfterTaxIncome"`
	FederalTaxesPaid float64 `json:"FederalTaxesPaid"`
	StateTaxesPaid   float64 `json:"StateTaxesPaid"`
	OtherTaxesPaid   float64 `json:"OtherTaxesPaid"`
}

func AfterTaxIncomeLambda(event Event) (MyResponse, error) {

	income := event.Income
	selfEmpStatus := event.SelfEmploymentStatus

	// Get the various tax values
	var federalTax float64 = FederalTaxes.Calculate(income)
	var stateTax float64 = StateTaxes.Calculate(income)
	var otherTax float64 = helpers.CalculateOtherTaxes(income, selfEmpStatus)
	afterTax := income - (federalTax + stateTax + otherTax)

	// create the output struct
	afterTaxOutput := MyResponse{
		AfterTaxIncome:   afterTax,
		FederalTaxesPaid: federalTax,
		StateTaxesPaid:   stateTax,
		OtherTaxesPaid:   otherTax,
	}
	return afterTaxOutput, nil
}
