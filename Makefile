
DATE = $(shell date "+%y.%m.%d.%s")

server.pb.go:
	protoc --go_out=plugins=grpc:. ./rpc/server.proto
pem:
	mkdir base/pem -p
	openssl genpkey -algorithm ED25519 -out base/pem/server.key &&openssl req -new -x509 -key base/pem/server.key -out base/pem/server.crt -subj "/CN=localhost" -days 3650
	openssl genpkey -algorithm ED25519 -out base/pem/client.key &&openssl req -new -x509 -key base/pem/client.key -out base/pem/client.crt -subj "/CN=localhost" -days 3650
# 	openssl genpkey -algorithm ED25519 -out client.key &&openssl req -new -key client.key -out client.csr -subj "/CN=localhost" && openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt

.PHONY:\
	pem \
	build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s" -o gohost .
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s" -o gohost-cgo .
ddoor: build
	ssh door "systemctl stop cj_hl.service"
	scp ./gohost  door:/root/hostl
	ssh door "systemctl restart cj_hl.service"
dqb: build
	ssh qb "systemctl stop cj_hl.service"
	scp ./gohost  qb:/root/hostl
	ssh qb "systemctl restart cj_hl.service"
dt:
	scp ./gohost.service t:/etc/systemd/system/
	GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
	ssh t "systemctl stop gohost.service"
	scp ./gohost t:/root/gohost
	ssh t "systemctl restart gohost.service"

vultr: build
	docker build . --tag sjc.vultrcr.com/cjie9759/gohost:$(DATE)
	docker build . --tag sjc.vultrcr.com/cjie9759/gohost:latest
	docker push sjc.vultrcr.com/cjie9759/gohost:$(DATE)
	docker push sjc.vultrcr.com/cjie9759/gohost:latest

	docker build . -f dockerfile-debian --tag sjc.vultrcr.com/cjie9759/gohost-debian:$(DATE)
	docker build . -f dockerfile-debian --tag sjc.vultrcr.com/cjie9759/gohost-debian:latest
	docker push sjc.vultrcr.com/cjie9759/gohost-debian:$(DATE)
	docker push sjc.vultrcr.com/cjie9759/gohost-debian:latest
docker: build
	docker build . --tag docker.io/cjie9759/gohost:$(DATE)
	docker build . --tag docker.io/cjie9759/gohost:latest
	docker push docker.io/cjie9759/gohost:$(DATE)
	docker push docker.io/cjie9759/gohost:latest
