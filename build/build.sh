set -e
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
ssh ten "systemctl stop cj_hs.service"
scp  ./hostListen ten:/root/hostls
ssh ten "systemctl restart cj_hs.service"

