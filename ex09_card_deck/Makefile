GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=deck


all: build test

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)

test:
	$(GOTEST) -v ./pkg/...

run:
	$(GORUN) ./cmd/main.go
