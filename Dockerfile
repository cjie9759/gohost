# https://github.com/librespeed/speedtest-go/releases/download/v1.1.4/speedtest-go_1.1.4_linux_amd64.tar.gz

FROM docker.io/alpine

WORKDIR /app

COPY gohost /app/gohost
ENTRYPOINT ["/app/gohost"]