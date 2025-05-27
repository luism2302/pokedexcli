build:
	@go build -o bin/pokedex/pokedex cmd/pokedex/main.go
run:
	@go run cmd/pokedex/main.go
test:
	@go test -v ./...