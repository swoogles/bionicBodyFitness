package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
	"net/url"
)

// TODO Get this fully reduced to the AWS request/response junk
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parameterMap, _ := url.ParseQuery(request.Body)
	token := parameterMap.Get("stripeToken")

	paymentResult := stripe.ExecuteTestStripePaymentWithAmount(token, 350)
	productId := "prod_Et1gMmK1DWlq3S"
	planId := stripe.CreateTestPlan(productId)
	customerId := stripe.CreateTestCustomer()
	stripe.CreateTestSubscription(planId, customerId)
	// TODO get some info about the charge, and then decide where to reroute

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       paymentResult,
	}, nil
}

func main() {
	lambda.Start(handler)
}
