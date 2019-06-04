GRPC_GOOGLE_APIS:=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
SRC:=$(GOPATH)/src
PROTO_PATH:=./proto/*.proto
DB_HOST		?= localhost
DB_NAME		?= ny_cab_data
DB_USER		?= root
DB_PWD		?= abcd1234

default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

install: ## build and install go application executable
	go install -v ./...

clean:  ## go clean
	go clean

tools: ## Fetch and install required tools
	go get -u github.com/grpc-ecosystem/grpc-gateway/...
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

compile-protobuf: ## Compile protocol buffer files
	protoc -I. -I$(GOPATH)/src \
	-I$(GRPC_GOOGLE_APIS) --go_out=plugins=grpc:. \
	$(PROTO_PATH)
	protoc -I. \
	-I$(SRC) \
	-I$(GRPC_GOOGLE_APIS) \
	--grpc-gateway_out=logtostderr=true:. \
	$(PROTO_PATH)

compile-binary: ## Create Linux ELF binary
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

generate-swagger: ## Generate swagger docs from protobuf files
	protoc -I. \
	-I$(GOPATH)/src \
	-I$(GRPC_GOOGLE_APIS) \
	--swagger_out=logtostderr=true:. \
	$(PROTO_PATH)

db-migrate:    ## Migrate DB schema. Set DB_PWD and DB_HOST to run against a specific environment. DB_USER and DB_NAME are set by convention.
#TODO: migrate cli might not exist. Need to have check & install step before running migrations
#See https://github.com/mattes/migrate/tree/master/cli
	# Starting migration
	@migrate -database 'mysql://${DB_USER}:${DB_PWD}@tcp(${DB_HOST}:3306)/${DB_NAME}?charset=utf8' -path  ./migrations/ up
	# Migration finished