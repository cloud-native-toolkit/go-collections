# Environment variables for the build process

PROJECTNAME=go-collections

MODULE=github.com/cloud-native-toolkit/go-collections

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE) # You can remove or change the path after last colon.
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
GOTESTS=$(shell find . -name "*_test.go" | grep -v vendor | xargs -I{} dirname {} | sed -E "s~[.]~${MODULE}~g")
VERSION ?= 0.1.0

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will store the server process id when it's running on development mode
PID=/tmp/.$(PROJECTNAME)-api-server.pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent


go-compile: go-clean go-get go-build

go-build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -ldflags "-X main.version=$(VERSION)" -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-generate:
	@echo "  >  Generating dependency files..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get)

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

go-test:
	@echo "  >  Testing code $(MODULE): $(GOTESTS)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GOTESTS)

## install: Install missing dependencies. Runs `go get` internally.
install: go-get

## build: Builds the application
build: go-build

## test: Tests the application
test: go-test

## compile: Compile the binary.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## clean: Clean build files. Runs `go clean` internally.
clean:
	@(MAKEFILE) go-clean

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
