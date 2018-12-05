all: pkg-test

pkg-test:
	@go test ./pkg/...