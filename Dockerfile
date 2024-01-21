# https://github.com/librespeed/speedtest-go/releases/download/v1.1.4/speedtest-go_1.1.4_linux_amd64.tar.gz

FROM docker.io/alpine

WORKDIR /app

COPY hostListen /app/hostListen
ENTRYPOINT ["/app/hostListen","-l","cjie.cf:17126"]