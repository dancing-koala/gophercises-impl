GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gophercise_1
BINARY_WIN=$(BINARY_NAME).exe

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

build_opt: 
	$(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)
	rm -f ./$(BINARY_WIN)

test:
	$(GOTEST) -v ./pkg/...

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o $(BINARY_WIN) -v