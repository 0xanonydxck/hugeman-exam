run:
	swag init && go run main.go
test:
	go test -v ./internal/service/...