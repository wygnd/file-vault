generate-go:
	mkdir -p ./file-service/gen
	protoc \
		--go_out=./file-service \
		--go_opt=module=github.com/wygnd/file-vault/file-service \
		--go-grpc_out=./file-service \
		--go-grpc_opt=module=github.com/wygnd/file-vault/file-service \
		shared/proto/file/file.proto

generate-nestjs:
	protoc --ts_out=./gateway/gen proto/file/file.proto

gen-all: generate-go generate-nestjs