package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
	"net/url"
)

// TODO Get this fully reduced to the AWS request/response junk
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parameterMap, _ := url.ParseQuery(request.Body)
	token := parameterMap.Get("stripeToken")
	fmt.Println(stripe.JsonSerialize(parameterMap))

	fmt.Println(stripe.JsonSerialize(request))
	//paymentResult := stripe.ExecuteTestStripePaymentWithAmount(token, 350)
	//productId := "prod_Et1gMmK1DWlq3S"
	//planId := stripe.CreateTestPlan(productId)
	//customerId := stripe.CreateTestCustomer(token)
	//subscription := stripe.CreateTestSubscription(planId, customerId)

	// TODO stripe.CreateOrder()
	result := stripe.ExecuteStripePaymentWithAmount(token, 123, "TEST_STRIPE_SECRET_KEY")
	// TODO get some info about the charge, and then decide where to reroute

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "new customer with order: " + result,
	}, nil
}

func main() {
	lambda.Start(handler)
}
