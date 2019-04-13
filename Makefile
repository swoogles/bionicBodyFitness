build:
	hugo -v
	mkdir -p functions
	# go get ./...
	go build -o lambdas/test-submit-stripe-charge  ./functions/stripe/test
	go build -o lambdas/live-submit-stripe-charge  ./functions/stripe/live
