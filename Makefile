obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu
receiver:
	@go run ./data_receiver
caculator:
	@go run ./distance_calculator

.PHONY: obu
