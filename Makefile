build:
	@go mod tidy
	@go build
clean:
	@rm -f ddoser
