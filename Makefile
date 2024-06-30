build:
	@go build -o ./bin/user_manager

run: build
	@./bin/user_manager

tests:
	@go test -v ./...