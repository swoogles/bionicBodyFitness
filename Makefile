build:
	mkdir -p lambdas
	go get ./...
	go build -o lambdas/test-charge  ./functions/stripe/test
	go build -o lambdas/live-submit-stripe-charge  ./functions/stripe/live
	go build -o lambdas/test-list-products  ./functions/stripe/products
	go build -o lambdas/after-login  ./functions/identity
	# Is this actually deploying?
	go build -o lambdas/register-new-member  ./functions/stripe/membership/new
	# hugo server --buildDrafts --buildFuture --ignoreCache
	hugo -v

