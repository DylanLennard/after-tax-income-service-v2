package main

import (
	"log"
	"net/http"

	"github.com/DylanLennard/after-tax-income-service-v2/service"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// setup the router
	r := httprouter.New()

	// assign routes
	r.GET("/", service.HelloWorld)
	r.GET("/after_tax_income", service.AfterTaxIncome)

	// start server and have router use defaultservermux
	log.Fatal(http.ListenAndServe(":8080", r))
}
