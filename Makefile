.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-gin-backend-admin

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf go-gin-backend-admin
	go clean -i .

swagger:
	swag init

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
