.PHONY: default run build test docs clean

APP_NAME=gopportunities
MAIN=cmd/main.go

default: run

run:
	@go run $(MAIN)
build:
	@go build -o $(APP_NAME) $(MAIN)
test:
	@go test ./ ...
docs:
	@swag init -g $(MAIN)
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs