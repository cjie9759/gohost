package user

import (
	"fmt"
	"hostListen/base"
	"hostListen/client"
	"log"
	"net/rpc"
	"time"
)

func showHostData() {
	for k, v := range base.HostData {
		fmt.Println("Host", k)
		for _, v1 := range v {
			fmt.Println()
			fmt.Println(v1.String())
		}
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
		log.Println("server return", result)
	}
	for {
		<-t.C
		f()
	}
}
