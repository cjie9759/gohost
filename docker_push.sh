tag=$(date "+%y.%m.%d.%s")
docker build . --tag docker.io/cjie9759/gohost:$tag
docker build . --tag docker.io/cjie9759/gohost:latest
docker push docker.io/cjie9759/gohost:$tag
docker push docker.io/cjie9759/gohost:latest
