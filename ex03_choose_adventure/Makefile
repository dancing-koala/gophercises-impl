GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=cyoa


build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)

run:
	$(GORUN) ./cmd/main.go
