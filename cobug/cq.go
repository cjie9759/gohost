package cobug

import "strconv"

type Cq struct {
	group_id string
	// message  string
	// auto_escape bool
	addr string
}

func NewCq(cq_addr string, GroupId int) *Cq {
	if cq_addr == "" {
		cq_addr = "http://127.0.0.1:5700"
	}
	return &Cq{
		addr: cq_addr + "/send_msg",
		// auto_escape: false,
		group_id: strconv.Itoa(GroupId),
	}
}
func (t *Cq) Send(msg string) {
	cqData := map[string]string{
		"group_id":    t.group_id,
		"message":     msg,
		"auto_escape": "false",
	}
	go NewBug().Post(t.addr, UrlEncode(cqData))
}
