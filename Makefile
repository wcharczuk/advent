all: pkg-test

profanity:
	@go run pkg/profanity/cmd/profanity.go

pkg-test:
	@go test ./pkg/...