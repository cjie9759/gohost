package user

import (
	"fmt"
	"gohost/base"
	"gohost/client"
	hostinfo "gohost/hostInfo"
	"log"
	"net/rpc"
	"sort"
	"time"
)

func showHostData(hs []hostinfo.HostInfo) {
	// fmt.Printf("\x1bc")
	fmt.Printf("\x1b[2J")
	sort.Slice(hs, func(i, j int) bool {
		return hs[i].Sid > hs[j].Sid
	})
	for _, k := range hs {
		fmt.Println(k.String())
	}
}
func User() {
	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()
	f := func() {
		conn, err := client.Con(base.Listen[0])
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		client := rpc.NewClient(conn)

		//调用方法
		hs := []hostinfo.HostInfo{}
		err = client.Call("Server.GetData", 1, &hs)
		showHostData(hs)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("server return")
	}
	for {
		<-t.C
		go f()
	}
}
