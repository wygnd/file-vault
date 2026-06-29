# FILE SERVICE (Go)
.PHONY: generate-go-file_service
generate-go-file_service:
	mkdir -p ./file-service/gen
	protoc \
		--go_out=./file-service \
		--go_opt=module=github.com/wygnd/file-vault/file-service \
		--go-grpc_out=./file-service \
		--go-grpc_opt=module=github.com/wygnd/file-vault/file-service \
		shared/proto/file/file.proto

# API GATEWAY (NestJS)
PROTO_SRC          := shared/proto
GATEWAY_GEN        := gateway/gen
PLUGIN             := gateway/node_modules/.bin/protoc-gen-ts_proto
PROTO_LIST         := $(shell find $(PROTO_SRC) -name "*.proto")

export PATH         := $(dir $(shell which node)):$(PATH)

.PHONY: generate-nestjs clean-generate
generate-nestjs: clean-generate
	mkdir -p $(GATEWAY_GEN)
	protoc \
		--plugin=$(PLUGIN) \
		--ts_proto_out=$(GATEWAY_GEN) \
		--ts_proto_opt=nestJs=true \
		--ts_proto_opt=outputEncodeMethods=false \
		--ts_proto_opt=outputJsonMethods=false \
		--ts_proto_opt=outputClientImpl=false \
		--ts_proto_opt=addGrpcMetadata=true \
		--ts_proto_opt=stringEnums=true \
		-I $(PROTO_SRC) \
		$(PROTO_LIST)

clean-generate:
	rm -rf $(GATEWAY_GEN)

# ALL
.PHONY: gen-all
generate-all: generate-go-file_service generate-nestjs