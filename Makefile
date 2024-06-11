server.pb.go:
	protoc --go_out=plugins=grpc:. ./rpc/server.proto
pem:
	mkdir base/pem -p
	openssl genpkey -algorithm ED25519 -out base/pem/server.key &&openssl req -new -x509 -key base/pem/server.key -out base/pem/server.crt -days 3650
	openssl genpkey -algorithm ED25519 -out base/pem/client.key &&openssl req -new -x509 -key base/pem/client.key -out base/pem/client.crt -days 3650

.PHONY:\
	pem \
	build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .