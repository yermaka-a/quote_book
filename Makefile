run: build
	@./bin/quote_book

build:
	@go build ./cmd/main.go -o bin/quoute_book .