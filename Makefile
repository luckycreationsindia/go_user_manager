build:
	@go build -o ./bin/user_manager

build-release:
	@go build -ldflags "-s -w" -o ./bin/release/user_manager

run: build
	@air

tests:
	@go test -v ./...