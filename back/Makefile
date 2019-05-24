UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GOOS ?= linux
endif
ifeq ($(UNAME_S),Darwin)
	GOOS ?= darwin
endif
VERSION=$(shell git describe --tags 2>/dev/null || echo 'Unknown')
BUILD_DATETIME=$(shell date '+%Y-%m-%d %H:%M:%S %Z')
GOCMD=go
GOBUILD=GOOS=$(GOOS) $(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=GOCACHE=off ginkgo
GOGET=$(GOCMD) get
LDFLAGS=-w -s -X 'main.version=$(VERSION)' -X 'main.buildDateTime=$(BUILD_DATETIME)'
BUILD_FLAGS=-a -installsuffix cgo --tags satic -ldflags "$(LDFLAGS)"

IMAGE_TAG=latest
TARGET=$(GOPATH)/src/master-your-life/target
PROJECT_PATH=$(GOPATH)/src/master-your-life

ACTIVITY_MANAGER_BINARY_NAME=activity-manager
ACTIVITY_PRODUCER_BINARY_NAME=activity-producer
ACTIVITY_PROVIDER_BINARY_NAME=activity-provider
GOALS_MANAGER_BINARY_NAME=goals-manager
USER_MANAGER_BINARY_NAME=user-manager

# Golang build commands
.PHONY: all
all: build

deps:
	dep ensure

.PHONY: build
build: activity-manager activity-producer activity-provider goals-manager user-manager

.PHONY: activity-manager
activity-manager:
	$(GOBUILD) $(BUILD_FLAGS) -o $(TARGET)/$(ACTIVITY_MANAGER_BINARY_NAME) activity-manager/main.go

.PHONY: activity-producer
activity-producer:
	$(GOBUILD) $(BUILD_FLAGS) -o $(TARGET)/$(ACTIVITY_PRODUCER_BINARY_NAME) activity-producer/main.go

.PHONY: activity-provider
activity-provider:
	$(GOBUILD) $(BUILD_FLAGS) -o $(TARGET)/$(ACTIVITY_PROVIDER_BINARY_NAME) activity-provider/main.go

.PHONY: goals-manager
goals-manager:
	$(GOBUILD) $(BUILD_FLAGS) -o $(TARGET)/$(GOALS_MANAGER_BINARY_NAME) goals-manager/main.go

.PHONY: user-manager
user-manager:
	$(GOBUILD) $(BUILD_FLAGS) -o $(TARGET)/$(USER_MANAGER_BINARY_NAME) user-manager/main.go

GIT_COMMIT	:=$(shell git rev-list -1 HEAD)
GIT_BRANCH	:=$(shell git rev-parse --abbrev-ref HEAD)
GO_VERSION 	:=$(shell go version)

LDFLAGS_STREAM_REPORT=-w -s -X '$(STREAM_REPORT_PACKAGE_PATH)/cmd.GitCommit=$(GIT_COMMIT)' -X '$(STREAM_REPORT_PACKAGE_PATH)/cmd.GitBranch=$(GIT_BRANCH)' -X '$(STREAM_REPORT_PACKAGE_PATH)/cmd.BuildDate=$(BUILD_DATETIME)' -X '$(STREAM_REPORT_PACKAGE_PATH)/cmd.GoVersion=$(GO_VERSION)' -X '$(STREAM_REPORT_PACKAGE_PATH)/cmd.Version=$(VERSION)'
STREAM_REPORT_BUILD_FLAGS = -a -installsuffix cgo --tags satic -ldflags "$(LDFLAGS_STREAM_REPORT)"


# Golang tests commands.
.PHONY: test
test: 	activity-manager-test \
		activity-producer-test \
		activity-provider-test \
		goals-manager-test \
		user-manager-test

.PHONY: activity-manager-test
activity-manager-test:
	GOOS=$(GOOS) go test ./activity-manager/...

.PHONY: activity-producer-test
activity-producer-test:
	GOOS=$(GOOS) go test ./activity-producer/...

.PHONY: activity-provider-test
activity-provider-test:
	GOOS=$(GOOS) go test ./activity-provider/...

.PHONY: goals-manager-test
goals-manager-test:
	GOOS=$(GOOS) go test ./goals-manager/...

.PHONY: user-manager-test
user-manager-test:
	GOOS=$(GOOS) go test ./user-manager/...