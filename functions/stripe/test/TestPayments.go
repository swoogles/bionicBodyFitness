package test

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"os"
)

func executeStripePayment(stripePaymentToken string) string {
	return executeStripePaymentWithAmount(stripePaymentToken, 999)
}

func executeStripePaymentWithAmount(stripePaymentToken string, amount int64) string {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Example charge"),
	}
	params.SetSource(stripePaymentToken)
	ch, _ := charge.New(params)
	fmt.Println(ch)

	return "faux payment response"
}
