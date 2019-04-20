package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	//m, _ := url.ParseQuery(request.Body)

	// Token is created using Checkout or Elements!
	// Get the payment token ID submitted by the form:

	stripe.GetAllTestProducts()

	fullLineupData := "Listed all products!"
	// TODO get some info about the charge, and then decide where to reroute

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fullLineupData,
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Set your secret key: remember to change this to your live secret key in production
// See your keys here: https://dashboard.stripe.com/account/apikeys
