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

	email := parameterMap.Get("email")
	name := parameterMap.Get("name")
	planId := parameterMap.Get("planId")

	var customerId string
	existingCustomerid := parameterMap.Get("stripe_customer_id")
	fmt.Println("existingCustomerid: " + existingCustomerid)
	if existingCustomerid == "undefined" || len(existingCustomerid) == 0 {
		customerId = stripe.CreateCustomer("STRIPE_SECRET_KEY", token, email, name)
		fmt.Println("New Customer Id: " + customerId)
	} else {
		customerId = existingCustomerid
	}

	stripe.CreateSubscription("STRIPE_SECRET_KEY", planId, customerId)

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
