all:
	@echo "Use a specific goal. To list all goals, type 'make help'"

.PHONY: dependencies # Lists project dependencies
dependencies:
	@gb vendor list

.PHONY: restore-dependencies # Restore project dependencies
restore-dependencies:
	@gb vendor restore

.PHONY: build # Builds binary executable
build:
	@gb build

.PHONY: docker-build # Builds Docker image
docker-build:
	@GOPATH=$(GOPATH):$(shell pwd)/vendor CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/cmd/main/
	@docker build -t segence/chartmuseum-authserver:$(shell make version) -f Dockerfile .

.PHONY: version # Prints project version
version:
	@cat VERSION

.PHONY: up # Runs Chartmuseum cluster
up:
	@VERSION=$(shell make version) docker-compose up -d

.PHONY: down # Tears down Chartmuseum cluster
down:
	@VERSION=$(shell make version) docker-compose down

.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1: \2/'
