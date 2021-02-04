
test:
	go test ./... -race -cover

demo:
	go run _examples/columns/main.go
