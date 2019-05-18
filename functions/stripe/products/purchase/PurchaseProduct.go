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
	productId := parameterMap.Get("planId") // TODO Fix name on front end

	var customerId string
	existingCustomerid := parameterMap.Get("stripe_customer_id")
	fmt.Println("existingCustomerid: " + existingCustomerid)
	if existingCustomerid == "undefined" || len(existingCustomerid) == 0 {
		customer, err := stripe.FindCustomer("STRIPE_SECRET_KEY", email)
		if err != nil {
			fmt.Println("Could not retrieve customer by email: " + err.Error())
			// TODO or should I fail here? maybe customer creation here gives us the best experience?
			customerId = stripe.CreateCustomer("STRIPE_SECRET_KEY", token, email, name)
		} else {
			customerId = customer.ID
			fmt.Println("Existing Customer Id retrieved by email: " + customerId)
		}
	} else {
		fmt.Println("Existing Customer Id attached to request: " + customerId)
		customerId = existingCustomerid
	}

	fmt.Println("OrderID: " + stripe.CreateOrder(productId, "STRIPE_SECRET_KEY", customerId))

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers: map[string]string{
			"location": "/punch_card_purchase_thanks",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}
