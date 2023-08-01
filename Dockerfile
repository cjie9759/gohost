# https://github.com/librespeed/speedtest-go/releases/download/v1.1.4/speedtest-go_1.1.4_linux_amd64.tar.gz

FROM alpine

WORKDIR /app

COPY hostListen /app/hostListen
RUN apk add --no-cache openssh-client

EXPOSE 80

ENTRYPOINT ["./hostListen -l cjie.cf:17126"]