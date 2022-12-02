package cq_test

import (
	"hostListen/pkg/cq"
	"testing"
)

func TestXxx(t *testing.T) {
	cq.Init("http://127.0.0.1:5700/send_msg", 938132468, true).Send("??????????????")
	cq.Cq.Send("test")
}
