package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"net/url"
	"os"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	m, _ := url.ParseQuery(request.Body)
	token := m.Get("stripeToken")

	stripe.Key = os.Getenv("LIVE_STRIPE_SECRET_KEY")

	// Token is created using Checkout or Elements!
	// Get the payment token ID submitted by the form:

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(828),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Example charge"),
	}
	params.SetSource(token)
	ch, _ := charge.New(params)
	fmt.Println(ch)

	fullLineupData := "faux payment response"
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
