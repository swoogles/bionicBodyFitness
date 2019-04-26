package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	billStripe "github.com/swoogles/stripe"
	"net/url"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	m, _ := url.ParseQuery(request.Body)

	fmt.Println("Request: ")
	fmt.Println(billStripe.JsonSerialize(request))
	fmt.Println("Request.Body: ")
	fmt.Println(billStripe.JsonSerialize(request.Body))
	fmt.Println("m: ")
	fmt.Println(billStripe.JsonSerialize(m))
	// Token is created using Checkout or Elements!
	// Get the payment token ID submitted by the form:

	// TODO get some info about the charge, and then decide where to reroute

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       billStripe.JsonSerialize(request),
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Set your secret key: remember to change this to your live secret key in production
// See your keys here: https://dashboard.stripe.com/account/apikeys
