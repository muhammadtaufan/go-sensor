build:
	go build -o bin/go-sensor cmd/main.go

run: build
	./bin/go-sensor

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/sensor.proto

.PHONY: proto
