package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	billStripe "github.com/swoogles/stripe"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Request: ")
	fmt.Println(billStripe.JsonSerialize(request))
	out, _ := json.Marshal(GetPunchCards())
	fmt.Println(string(out))
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(out),
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Set your secret key: remember to change this to your live secret key in production
// See your keys here: https://dashboard.stripe.com/account/apikeys

func GetPunchCards() []billStripe.Product {
	return billStripe.GetActiveProductsWithUnsafeType("STRIPE_SECRET_KEY", string(stripe.ProductTypeGood))
}
