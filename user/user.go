package user

import (
	"fmt"
	"hostListen/base"
	"hostListen/client"
	"log"
	"net/rpc"
	"sort"
	"time"
)

func showHostData() {
	// fmt.Printf("\x1bc")
	fmt.Printf("\x1b[2J")
	ks := make([]string, 0, len(base.HostData))
	for k := range base.HostData {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		v := base.HostData[k]
		v1 := v[len(v)-1]
		fmt.Println(v1.String())
	}
}
func User() {
	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()
	f := func() {
		conn, err := client.Con()
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		client := rpc.NewClient(conn)

		//调用方法
		result := base.HostData
		err = client.Call("Server.GetData", 1, &result)
		showHostData()
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
