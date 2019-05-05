package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	billStripe "github.com/swoogles/stripe"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Request: ")
	fmt.Println(billStripe.JsonSerialize(request))
	out, _ := json.Marshal(billStripe.GetAllPlans("STRIPE_SECRET_KEY"))
	fmt.Println(string(out))
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(out),
	}, nil
}

func main() {
	lambda.Start(handler)
}
