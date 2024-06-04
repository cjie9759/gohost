server.pb.go:
	protoc --go_out=plugins=grpc:. ./rpc/server.proto