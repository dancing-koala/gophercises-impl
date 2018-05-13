GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=task


build: 
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

clean:
	$(GOCLEAN)
	rm -f ./$(BINARY_NAME)

run:
	make build
	rm ./db/gophercise_7.db
	./$(BINARY_NAME) add task 1
	./$(BINARY_NAME) add task 2
	./$(BINARY_NAME) list
	./$(BINARY_NAME) do 1
	./$(BINARY_NAME) list
	./$(BINARY_NAME) completed
	./$(BINARY_NAME) rm 1
	./$(BINARY_NAME) completed
