
test:
	go test ./... -race -cover -v

demo:
	go run _examples/columns/main.go
