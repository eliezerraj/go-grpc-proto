# Define environment variables
GO_MODULE := github.com/eliezerraj/go-grpc-proto

# Default target
all: protoc

clean:
	@echo "Cleaning protogen files..."
	rm -fR ./protogen 
	mkdir -p ./protogen

protoc:
	@echo "Compile protoc files..."
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/token/payment/*.proto \
	./proto/token/health/*.proto \
	./proto/token/pod/*.proto \
	./proto/token/card/*.proto \
	./proto/token/*.proto 

build: clean protoc

pipeline-init:
	@echo "Running pipeline-init ..."
	sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

pipeline-build:	pipeline-init build

.PHONY: all clean protoc build pipeline-init pipeline-build