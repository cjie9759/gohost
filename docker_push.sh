set -e
tag=$(date "+%y.%m.%d.%s")
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
podman build . --tag docker.io/cjie9759/gohost:$tag
podman build . --tag docker.io/cjie9759/gohost:latest
podman push docker.io/cjie9759/gohost:$tag
podman push docker.io/cjie9759/gohost:latest
