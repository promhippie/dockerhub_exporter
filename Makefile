PROMU := $(GOPATH)/bin/promu
PKGS := $(shell go list ./... | grep -v /vendor/)

PREFIX ?= $(shell pwd)
BIN_DIR ?= $(shell pwd)
DOCKER_IMAGE_NAME ?= dockerhub-exporter
DOCKER_IMAGE_TAG ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

all: format build test

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

test:
	@echo ">> running tests"
	@go test -short $(PKGS)

format:
	@echo ">> formatting code"
	@go fmt $(PKGS)

vet:
	@echo ">> vetting code"
	@go vet $(PKGS)

build: promu
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball $(BIN_DIR) --prefix $(PREFIX)

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

promu:
	@which promu > /dev/null; if [ $$? -ne 0 ]; then \
		@go get -u github.com/prometheus/promu; \
	fi

.PHONY: all style format build test vet tarball docker promu
