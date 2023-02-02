NAME    = zyh-webserver
VERSION = 0.0.1
BUILD   = ./build

.PHONY: docker
docker:
	@docker-compose -p "zyh-gin-webapp" up --build --detach

.PHONY: docker-down
docker-down:
	@docker-compose -p "zyh-gin-webapp" down
## build: Compile the packages.
.PHONY: build
build:
	@go build -o $(BUILD)/$(NAME) .

## run: Build and Run in development mode.
.PHONY: run
run: build
	@$(BUILD)/$(NAME) -env local

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@$(BUILD)/$(NAME) -env production

## test: Run tests with verbose mode
.PHONY: test
test:
	@go test -v ./test

## test with coverage
.PHONY: test-cov
test-cov:
	@go test ./test -coverpkg=./src/... -race -covermode=atomic -coverprofile=coverage.out

## deps: Download modules
.PHONY: deps
deps:
	@go mod download

## clean: Clean project and previous builds.
.PHONY: clean
clean:
	@rm -rf $(BUILD) *.out

########## Following are unverified scripts




.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo