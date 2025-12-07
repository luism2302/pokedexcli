BINARY_NAME=pokedexcli

build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ./cmd/pokedexcli/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows ./cmd/pokedexcli/main.go
