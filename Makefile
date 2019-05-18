build:
	mkdir -p lambdas
	go get ./...
	go build -o lambdas/list-products  ./functions/stripe/products
	go build -o lambdas/purchase-product  ./functions/stripe/products/purchase
	go build -o lambdas/list-plans  ./functions/stripe/plans
	go build -o lambdas/after-login  ./functions/identity
	go build -o lambdas/find-stripe-customer  ./functions/identity/find_customer
	go build -o lambdas/create-stripe-customer  ./functions/stripe/customers/create
	go build -o lambdas/register-new-member  ./functions/stripe/membership/new
	# hugo server --buildDrafts --buildFuture --ignoreCache
	hugo -v

