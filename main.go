package main

import (
	"fmt"
	"github.com/DylanLennard/after-tax-service-v2/service"
)

func main() {
	/*
		Here's what we need to do
		* Make this effective CLI first
		* Then make this http handy
		* Then build unit tests
		* Then deploy this via lambda
		* then set up with API gateway if need endpoint
	*/
	fmt.Println(service.AfterTaxIncome(90000, true))
}
