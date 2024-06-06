server.pb.go:
	protoc --go_out=plugins=grpc:. ./rpc/server.proto
pem:
	mkdir pem -p
	openssl genpkey -algorithm ED25519 -out pem/server.key &&openssl req -new -x509 -key pem/server.key -out pem/server.crt -days 3650
	openssl genpkey -algorithm ED25519 -out pem/client.key &&openssl req -new -x509 -key pem/client.key -out pem/client.crt -days 3650

.PHONY:\
	pem
