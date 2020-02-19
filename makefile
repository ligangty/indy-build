# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGETU=$(GOGET) -u -v
BUILD_DIR=./build
BINARY_NAME=$(BUILD_DIR)/indy-build
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
		$(GOBUILD) -trimpath -o $(BINARY_NAME) -v ./cmd/indy-build
.PHONY: build

test: 
		$(GOTEST) -v ./...
.PHONY: test

clean: 
		$(GOCLEAN) ./...
		rm -rf $(BUILD_DIR)
.PHONY: clean
