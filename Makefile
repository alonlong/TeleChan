GO=GO111MODULE=on go
GOBUILD=$(GO) build
GOINSTALL=$(GO) install

default: build

all: build install proto

build:
	$(GOBUILD) -o TeleChan main.go

install:
	$(GOINSTALL)

proto:
	protoc --go_out=plugins=grpc:. ./pkg/api/api.proto
