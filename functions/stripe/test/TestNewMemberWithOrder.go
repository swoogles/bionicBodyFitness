package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
	"net/url"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parameterMap, _ := url.ParseQuery(request.Body)
	token := parameterMap.Get("stripeToken")
	fmt.Println("Request after stripe token replacement: " + stripe.JsonSerialize(request))

	customerId := stripe.CreateCustomer("TEST_STRIPE_SECRET_KEY", token, parameterMap.Get("emailaddress"), "name")
	paymentId := stripe.CreateOrder("sku_EuxgNOUYjIEyFZ", "TEST_STRIPE_SECRET_KEY", customerId)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "new customer with order: " + paymentId,
	}, nil
}

func main() {
	lambda.Start(handler)
}
