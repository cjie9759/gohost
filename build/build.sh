set -e
ssh ten "systemctl stop cj_hs.service"
scp  ./hostListen ten:/root/hostls
ssh ten "systemctl restart cj_hs.service"

