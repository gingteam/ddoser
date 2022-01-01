build:
	@go mod tidy
	@go build -ldflags="-s -w"
clean:
	@rm -f ddoser
