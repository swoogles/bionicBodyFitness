package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	billStripe "github.com/swoogles/stripe"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	email := request.QueryStringParameters["email"]
	x := "hi"
	fmt.Println(x)

	customer, err := billStripe.FindCustomer("STRIPE_SECRET_KEY", email)
	if err != nil {
		fmt.Println("Customer lookup error: " + err.Error())
	} else {
		fmt.Println(billStripe.JsonSerialize(customer))
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       billStripe.JsonSerialize(customer),
	}, nil
}

func main() {
	lambda.Start(handler)
}
