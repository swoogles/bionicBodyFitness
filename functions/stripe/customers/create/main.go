package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
	"net/http"
	"net/url"
)

// TODO Get this fully reduced to the AWS request/response junk
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parameterMap, _ := url.ParseQuery(request.Body)
	token := parameterMap.Get("stripeToken")
	//fmt.Println(stripe.JsonSerialize(parameterMap))

	fmt.Println("Request after stripe token replacement: " + stripe.JsonSerialize(request))

	customerId := stripe.CreateCustomer("STRIPE_SECRET_KEY", token, parameterMap.Get("emailaddress"), parameterMap.Get("name"))
	fmt.Println("New customerId: " + customerId)

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
