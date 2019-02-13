package service

import (
	"strconv"

	"github.com/DylanLennard/after-tax-income-service-v2/helpers"
)

// Event used for accepting parameters from querystring
type Event struct {
	Income               string `json:"Income"`
	SelfEmploymentStatus string `json:"SelfEmploymentStatus"`
}

// MyResponse used for setting up exported JSON after handler runs
type MyResponse struct {
	AfterTaxIncome   float64 `json:"AfterTaxIncome"`
	FederalTaxesPaid float64 `json:"FederalTaxesPaid"`
	StateTaxesPaid   float64 `json:"StateTaxesPaid"`
	OtherTaxesPaid   float64 `json:"OtherTaxesPaid"`
}

// AfterTaxIncomeLambda will calculate your after tax income based on
// the income provided and the self employment status pased as
// query parameters to the Event
// eventually this will be expanded to include State and there will
// be a data store with state data in there
// TODO: set default values for status and proper error handling
func AfterTaxIncomeLambda(event Event) (MyResponse, error) {

	incomeStr := event.Income
	selfEmpStatusStr := event.SelfEmploymentStatus

	// check for errors
	income, err := strconv.ParseFloat(incomeStr, 64)
	if err != nil {
		return MyResponse{}, err
	}
	selfEmpStatus, errEmpStatus := strconv.ParseBool(selfEmpStatusStr)
	if errEmpStatus != nil || selfEmpStatus {
		selfEmpStatus = false
	}

	// Get the various tax values
	federalTax := FederalTaxes.Calculate(income)
	stateTax := StateTaxes.Calculate(income)
	otherTax := helpers.CalculateOtherTaxes(income, selfEmpStatus)
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
