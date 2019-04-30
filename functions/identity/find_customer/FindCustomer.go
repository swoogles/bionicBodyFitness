package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	billStripe "github.com/swoogles/stripe"
	"net/url"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	parameterMap, _ := url.ParseQuery(request.Body)
	email := parameterMap.Get("email")
	if email == "" {
		email = request.PathParameters["email"]
	}

	fmt.Println("About to lookup Stripe customer. email: " + email)
	customer, err := billStripe.FindCustomer("TEST_STRIPE_SECRET_KEY", email)
	if err != nil {
		fmt.Println("Customer lookup error: " + err.Error())
	} else {
		fmt.Println(customer)
	}

	billStripe.JsonSerialize(request.PathParameters)
	//m, _ := url.ParseQuery(request.Body)

	fmt.Println("Request: ")
	fmt.Println(billStripe.JsonSerialize(request))
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       billStripe.JsonSerialize(customer),
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Set your secret key: remember to change this to your live secret key in production
// See your keys here: https://dashboard.stripe.com/account/apikeys
