GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=shortener


build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)

run:
	$(GORUN) ./cmd/main.go
