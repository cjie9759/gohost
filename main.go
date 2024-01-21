package main

import (
	"hostListen/base"
	"hostListen/client"
	"hostListen/server"
	"hostListen/user"
)

func main() {
	base.Init()

	switch {
	case base.Is_server:
		go server.Listen()
		server.TlsService()
	case base.Is_user:
		user.User()
	default:
		client.Client()
	}

	// 客户端
	// <-t.C
	// t := time.NewTicker(time.Minute / 10)
	// defer t.Stop()
	// for {
	// }

}
