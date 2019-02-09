package main

import (
	"github.com/DylanLennard/after-tax-service-v2/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func main() {
	/*
		Here's what we need to do
		* Build unit tests
		* Then deploy this via lambda
		* then set up with API gateway if need endpoint
	*/
	// setup the router
	r := httprouter.New()

	// assign routes 
	r.GET("/", service.HelloWorld)
	r.GET("/after_tax_income", service.AfterTaxIncome)

	// start server and have router use defaultservermux
	log.Fatal(http.ListenAndServe(":8080", r))
}
