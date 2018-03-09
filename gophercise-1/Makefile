GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gophercise_1

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

build_opt: 
	$(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)

test:
	$(GOTEST) -v ./pkg/...

run:
	$(GORUN) ./cmd/main.go
