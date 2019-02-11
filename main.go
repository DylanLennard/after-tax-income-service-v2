package main

import (
	"log"
	"net/http"

	"github.com/DylanLennard/after-tax-income-service-v2/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	/*
		Here's what we need to do
		* deploy this via lambda
		* set up with API gateway if need endpoint
		* after that, rewrite to work with docker
	*/
	// setup the router
	r := httprouter.New()

	// assign routes
	r.GET("/", service.HelloWorld)
	r.GET("/after_tax_income", service.AfterTaxIncome)

	// start server and have router use defaultservermux
	log.Fatal(http.ListenAndServe(":8080", r))
}
