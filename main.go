package main

import (
	"github.com/DylanLennard/after-tax-income-service-v2/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(service.AfterTaxIncomeLambda)
}
