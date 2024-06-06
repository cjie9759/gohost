package client

import (
	"gohost/base"
	hostinfo "gohost/hostInfo"
	"gohost/server"
	"log"
	"time"
)

func NewClien(s server.ServerInterface) *clientv2 {
	return &clientv2{s}
}

type clientv2 struct {
	s server.ServerInterface
}

func (c *clientv2) Run() {
	if base.Name != "" {
		base.Name = base.Name + "----"
	}

	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()

	var err error
	for {
		<-t.C
		err = c.s.Save(hostinfo.GetHostInfo())
		if err != nil {
			log.Println(err)
		}
	}

}
