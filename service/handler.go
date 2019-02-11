package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DylanLennard/after-tax-income-service-v2/helpers"
	"github.com/julienschmidt/httprouter"
)

type afterTaxStruct struct {
	AfterTaxIncome   float64
	FederalTaxesPaid float64
	StateTaxesPaid   float64
	OtherTaxesPaid   float64
}

// HelloWorld to test that the server is up running and working
func HelloWorld(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Hello Homie")
}

// This handler results in the calculation of your after tax income
// assuming you live in the state of california for the 2018 tax year.
// It'll eventually take in the URL and parse out your income and status level
// and pass that into the Tax function defined in helpers
// AFter that it calculates federal, state, and other taxes
// and subtracts them from your income and returns the final income
func AfterTaxIncome(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	queryVals := req.URL.Query()

	// get income
	incomeStr := queryVals.Get("income")
	income, err := strconv.ParseFloat(incomeStr, 64)
	if err != nil {
		fmt.Fprint(w, "Cannot process income parameter: ", incomeStr)
		return
	}

	// get employment status
	selfEmpStatusStr := queryVals.Get("status")
	selfEmpStatus, err_new := strconv.ParseBool(selfEmpStatusStr)
	if err_new != nil {
		selfEmpStatus = false
	}

	// Get the various tax values
	var federalTax float64 = FederalTaxes.Calculate(income)
	var stateTax float64 = StateTaxes.Calculate(income)
	var otherTax float64 = helpers.CalculateOtherTaxes(income, selfEmpStatus)
	afterTax := income - (federalTax + stateTax + otherTax)

	// create the output struct
	afterTaxOutput := afterTaxStruct{
		AfterTaxIncome:   afterTax,
		FederalTaxesPaid: federalTax,
		StateTaxesPaid:   stateTax,
		OtherTaxesPaid:   otherTax,
	}

	// marshal into JSON and write to response
	data, err := json.Marshal(&afterTaxOutput)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal("error:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}
