package main

import (
	"gohost/base"
	"gohost/client"
	"gohost/server"
	"gohost/user"
)

func main() {
	base.Init()

	switch {
	case base.Is_server:
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
