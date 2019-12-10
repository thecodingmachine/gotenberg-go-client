GOLANG_VERSION=1.13
GOTENBERG_VERSION=6
GOTENBERG_LOG_LEVEL=ERROR
VERSION=snapshot
GOLANGCI_LINT_VERSION=1.20.1

# gofmt and goimports all go files.
fmt:
	go fmt ./...
	go mod tidy

# run all linters.
lint:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) -t thecodingmachine/gotenberg-go-client:lint -f build/lint/Dockerfile .
	docker run --rm -it -v "$(PWD):/lint" thecodingmachine/gotenberg-go-client:lint

# run all tests.
tests:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOTENBERG_VERSION=$(GOTENBERG_VERSION) --build-arg GOTENBERG_LOG_LEVEL=$(GOTENBERG_LOG_LEVEL) -t thecodingmachine/gotenberg-go-client:tests -f build/tests/Dockerfile .
	docker run --rm -it -v "$(PWD):/tests" thecodingmachine/gotenberg-go-client:tests