TAG=feature-flag-service:latest

dev:
	go run main.go

test:
	go test -cover ./...

build:
	docker build -t $(TAG) .
