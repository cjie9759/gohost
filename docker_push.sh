set -e
tag=$(date "+%y.%m.%d.%s")
echo $tag
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
docker build . --tag docker.io/cjie9759/gohost:$tag
docker build . --tag docker.io/cjie9759/gohost:latest
docker push docker.io/cjie9759/gohost:$tag
docker push docker.io/cjie9759/gohost:latest


# docker run --hostname cc --restart=always --name gohost docker.io/cjie9759/gohost:latest
# docker run --hostname 43 --restart=always --name gohost docker.io/cjie9759/gohost:latest
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
podman build . --tag docker.io/cjie9759/gohost:$tag
podman build . --tag docker.io/cjie9759/gohost:latest
podman push docker.io/cjie9759/gohost:$tag
podman push docker.io/cjie9759/gohost:latest
# podman run --hostname cc --restart=always --name gohost docker.io/cjie9759/gohost:latest
# podman run --hostname 43 --restart=always --name gohost docker.io/cjie9759/gohost:latest