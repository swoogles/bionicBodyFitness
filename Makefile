build:
	hugo -v
	mkdir -p lambdas
	go get ./...
	go build -o lambdas/test-submit-stripe-charge  ./functions/stripe/test
	go build -o lambdas/live-submit-stripe-charge  ./functions/stripe/live
	go build -o lambdas/test-list-products  ./functions/stripe/products
