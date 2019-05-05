package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
	"net/http"
	"net/url"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parameterMap, _ := url.ParseQuery(request.Body)
	token := parameterMap.Get("stripeToken")

	fmt.Println("Request after stripe token replacement: " + stripe.JsonSerialize(request))

	customerId := stripe.CreateCustomerAndSubscribeToPlan(
		"STRIPE_SECRET_KEY",
		token,
		parameterMap.Get("email"),
		parameterMap.Get("name"),
		parameterMap.Get("planId"),
	)
	// TODO Get planId and create a subscription with it.
	fmt.Println(customerId)
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers: map[string]string{
			"location": "/register_new_member_thanks",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}
