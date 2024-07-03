build:
	@go build -o ./bin/user_manager

run: build
	@air

tests:
	@go test -v ./...