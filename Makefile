build:
	go build -o bin/quant main.go
.PHONY: build

run:
	go run main.go
.PHONY: run
