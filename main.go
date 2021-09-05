package main

import (
	"hostListen/base"
	"hostListen/client"
	"hostListen/server"
	"hostListen/user"
	"time"
)

func main() {
	if *base.Is_server {
		server.Service()
		return
	}

	if *base.Is_user {
		user.User()
		return
	}

	// 客户端
	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()
	for {
		client.Client()
		<-t.C
	}

}
