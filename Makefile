all : lint install

lint:
	golangci-lint run --fix

install:
	go install .
