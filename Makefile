all: build

.PHONY: build clean install

CURRENT_GOOS := $(shell go env GOOS)
CURRENT_GOARCH := $(shell go env GOARCH)

OUTPUT_FILE := oag$(if $(filter windows,$(CURRENT_GOOS)),.exe,)

build: deps
	go build -o $(OUTPUT_FILE) ./cmd/oag

deps:
	go mod tidy

install: build
	go install ./cmd/$(OUTPUT_FILE)

clean:
	@echo "Cleaning build directory..."
	@rm -f $(OUTPUT_FILE)
