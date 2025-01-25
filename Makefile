# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

# Binary names
BINARY_NAME=ethereum-wallet

all: build

build:
	env GO111MODULE=on $(GOBUILD) 

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	env GO111MODULE=on $(GOBUILD) 
	./$(BINARY_NAME)

fmt:
	$(GOFMT) ./...

vet:
	$(GOCMD) vet ./...

lint:
	golint ./...

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux -v ./...

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows.exe -v ./...

.PHONY: all build test clean run fmt vet lint build-linux build-windows
