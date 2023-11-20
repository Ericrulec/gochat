generate:
	protoc --proto_path=proto/gochat --go_out=. --go_opt=module=github.com/Ericrulec/gochat --go-grpc_out=. --go-grpc_opt=module=github.com/Ericrulec/gochat gochat.proto
