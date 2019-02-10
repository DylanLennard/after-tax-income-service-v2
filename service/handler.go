package service 

import (
	"fmt"
	"github.com/DylanLennard/after-tax-service-v2/helpers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// HelloWorld to test that the server is up running and working
func HelloWorld(w http.ResponseWriter, req *http.Request, params httprouter.Params){
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

	// income 
	incomeStr := queryVals.Get("income")
	income, err := strconv.ParseFloat(incomeStr, 64)
	if err != nil { 
		fmt.Fprint(w, "Cannot process income parameter: ", incomeStr)
		return
	}

	// employment status
	selfEmpStatusStr := queryVals.Get("status")
	selfEmpStatus, err_new := strconv.ParseBool(selfEmpStatusStr)
	if err_new != nil{
		selfEmpStatus = false
	}

	// Get the various tax values
	federalTax := FederalTaxes.Calculate(income)
	stateTax := StateTaxes.Calculate(income)
	otherTax := helpers.CalculateOtherTaxes(income, selfEmpStatus)
	afterTax := income - (federalTax + stateTax + otherTax)
	
	// write to response
	fmt.Fprintf(w, "%.2f", afterTax)
	
}
