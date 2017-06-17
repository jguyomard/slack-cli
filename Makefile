configFile=$(shell pwd)/config.yaml

help:
	@echo "Available commands:"
	@echo " - make build  - compile project and dependencies"
	@echo " - make lint   - run linter"
	@echo " - make test   - run all tests"
	@echo " - make clean  - remove bin files"

build:
	go get -d -v ./...
	go build -o slack-cli main.go

lint:
	go get -u github.com/golang/lint/golint
	golint ./...

test:
	go test ./src/commands/ -config-file=$(configFile) ${ARGS}

clean:
	rm -f slack-cli
