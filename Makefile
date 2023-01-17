NAME=go-webapp-gin
VERSION=0.0.1

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test -v ./src/...

.PHONY: test-cov
## test with coverage
test-cov:
	@go test ./src/*** -coverpkg=./src/... -race -covermode=atomic -coverprofile=coverage.out

########## Following are unverified scripts
.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e local

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download



.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo