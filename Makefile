SOURCES = $(sort $(dir $(wildcard ./api/proto/api/*/)))

.ONESHELL:

build-go:
	go build -v ./cmd/app

generate-pb:
	protoc -I ./api/proto --go_out ./extra --go_opt paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./extra --go-grpc_opt paths=source_relative --grpc-gateway_out ./extra --grpc-gateway_opt paths=source_relative --openapiv2_out ./docs --openapiv2_opt logtostderr=true ./api/proto/api/*.proto
	protoc-go-inject-tag -input="./extra/api/*"
evans:
	evans api/protofile/$(name).proto -p $(port)
generate-docs:
	redoc-cli bundle -o ./docs/index.html docs/api_docs.swagger.json