# set -e
# # CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
# GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s"  .
# ssh ten "systemctl stop cj_hs.service"
# scp ./hostListen ten:/root/hostls
# ssh ten "systemctl restart cj_hs.service"


ssh door "systemctl stop cj_hl.service"
scp ./gohost  door:/root/hostl
ssh door "systemctl restart cj_hl.service"

ssh qb "systemctl stop cj_hl.service"
scp ./gohost  qb:/root/hostl
ssh qb "systemctl restart cj_hl.service"

# ssh ten "docker stop gohost"
# ssh ten "docker rm gohost"
# ssh ten "docker pull docker.io/cjie9759/gohost:latest"
# ssh ten "docker run -d --name gohost --hostname t6 --restart=always docker.io/cjie9759/gohost:latest"

# docker stop gohost
# docker rm gohost
# docker pull docker.io/cjie9759/gohost:latest
# docker run -d --name gohost --hostname ten --restart=always docker.io/cjie9759/gohost:latest
