SOURCES = $(sort $(dir $(wildcard ./api/proto/api/*/)))

.ONESHELL:

build-go:
	go build -v ./cmd/app

generate-pb:
	 protoc -I ./api/proto --go_out ./extra --go_opt paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./extra --go-grpc_opt paths=source_relative --grpc-gateway _out ./extra --grpc-gateway_opt paths=source_relative ./api/proto/api/*.proto;
evans:
	evans api/protofile/$(name).proto -p $(port)

generate-docs:
	protoc -I ./api/proto --swagger_out=allow_merge=true,merge_file_name=api_docs:./docs/ \
                       ./api/proto/api/*;
	redoc-cli bundle -o ./docs/index.html docs/api_docs.swagger.json;
	rm ./docs/api_docs.swagger.json;