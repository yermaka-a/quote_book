run: build
	@./bin/quote_book

build:
	@go build -o  bin/quote_book ./cmd/main.go