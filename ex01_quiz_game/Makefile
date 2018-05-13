GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=quizz

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)

test:
	$(GOTEST) -v ./pkg/...

run:
	$(GORUN) ./cmd/main.go

run_alt:
	$(GORUN) ./cmd/main.go --data-path="./problems_alt.csv"

run_shuffle:
	$(GORUN) ./cmd/main.go -shuffle
