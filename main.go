package main

import (
	"hostListen/base"
	"hostListen/client"
	"hostListen/server"
	"hostListen/user"
)

func main() {
	if *base.Is_server {
		server.TlsService()
		return
	}

	if *base.Is_user {
		user.User()
		return
	}

	// 客户端
	client.Client()
	// <-t.C
	// t := time.NewTicker(time.Minute / 10)
	// defer t.Stop()
	// for {
	// }

}
