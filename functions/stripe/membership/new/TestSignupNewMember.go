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

	customerId := stripe.CreateCustomer("TEST_STRIPE_SECRET_KEY", token, parameterMap.Get("emailaddress"))
	planId := "plan_EuxyCM45UamtZl" // TODO Serve plans on page and take as param here

	stripe.CreateSubscription("TEST_STRIPE_SECRET_KEY", planId, customerId)

	//return &events.APIGatewayProxyResponse{
	//	StatusCode: 303,
	//	Body:       "new membership subscription id: " + subscriptionId,
	//}, nil

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusPermanentRedirect,
		Headers: map[string]string{
			"location": "/contact_thanks",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}
