PROTOC_VERSION=3.13.0
PROTOC_RELEASE_URL=https://github.com/protocolbuffers/protobuf/releases/download
ifeq "$(OS)" "Windows_NT"
	PROTOC?=protoc	
else
	PROTOC?=bin/protoc
    UNAME=$(shell uname -s)
	ifeq "$(UNAME)" "Linux"
		PROTOC_PKG=protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	else
		PROTOC_PKG=protoc-$(PROTOC_VERSION)-osx-x86_64.zip
	endif
endif
PROTOC_PKG_URL=$(PROTOC_RELEASE_URL)/v$(PROTOC_VERSION)/$(PROTOC_PKG)

# Path to tools
PROTOC_GEN_GO=bin/protoc-gen-go
PROTOC_GEN_GO_GRPC=bin/protoc-gen-go-grpc

.DEFAULT_GOAL=build

tmp bin:
	mkdir $@

tmp/$(PROTOC_PKG): tmp bin
	wget -q -O $@ $(PROTOC_PKG_URL)

$(PROTOC): tmp/$(PROTOC_PKG)
	unzip -q -o $<
	@rm readme.txt

.PHONY: protoc
protoc: $(PROTOC)

hello/pb/hello.pb.go: proto/hello.proto
	$(PROTOC) \
	-I $(dir $<) \
	--plugin protoc-gen-go=$(PROTOC_GEN_GO) \
	--plugin protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC) \
	--go-grpc_opt paths=source_relative \
	--go-grpc_out $(dir $@) \
	--go_opt paths=source_relative \
	--go_out $(dir $@) \
	$<

bin/hello-client: pb $(wildcard **/*.go)
	go build -o $@ ./cmd/client

bin/hello-server: pb $(wildcard **/*.go)
	go build -o $@ ./cmd/server

.PHONY: pb
pb: hello/pb/hello.pb.go

.PHONY: build
build: bin/hello-client bin/hello-server
