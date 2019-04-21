build:
	mkdir -p lambdas
	go get ./...
	go build -o lambdas/test-charge  ./functions/stripe/test
	go build -o lambdas/live-submit-stripe-charge  ./functions/stripe/live
	go build -o lambdas/test-list-products  ./functions/stripe/products
	# hugo server --buildDrafts --buildFuture --ignoreCache
	hugo -v

