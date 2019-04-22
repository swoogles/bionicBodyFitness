package new

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/stripe"
	"net/url"
)

// TODO Get this fully reduced to the AWS request/response junk
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parameterMap, _ := url.ParseQuery(request.Body)
	token := parameterMap.Get("stripeToken")
	//fmt.Println(stripe.JsonSerialize(parameterMap))

	fmt.Println("Request after stripe token replacement: " + stripe.JsonSerialize(request))
	//paymentResult := stripe.ExecuteTestStripePaymentWithAmount(token, 350)
	//productId := "prod_Et1gMmK1DWlq3S"
	//planId := stripe.CreateTestPlan(productId)
	//customerId := stripe.CreateTestCustomer(token)
	//subscription := stripe.CreateTestSubscription(planId, customerId)

	customerId := stripe.CreateCustomer("TEST_STRIPE_SECRET_KEY", token, parameterMap.Get("emailaddress"))
	planId := "plan_EuxyCM45UamtZl" // TODO Serve plans on page and take as param here
	subscriptionId := stripe.CreateSubscription("TEST_STRIPE_SECRET_KEY", planId, customerId)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "new membership subscription id: " + subscriptionId,
	}, nil
}

func main() {
	lambda.Start(handler)
}
