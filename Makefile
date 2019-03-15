GOLANG_VERSION=1.12
GOTENBERG_VERSION=4.2.1
VERSION=snapshot

# gofmt and goimports all go files.
fmt:
	go fmt ./...
	go mod tidy

# run all linters.
lint:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) -t thecodingmachine/gotenberg-go-client:lint -f build/lint/Dockerfile .
	docker run --rm -it -v "$(PWD):/lint" thecodingmachine/gotenberg-go-client:lint

# run all tests.
tests:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOTENBERG_VERSION=$(GOTENBERG_VERSION)  -t thecodingmachine/gotenberg-go-client:tests -f build/tests/Dockerfile .
	docker run --rm -it -v "$(PWD):/tests" thecodingmachine/gotenberg-go-client:tests