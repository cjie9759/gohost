package main

import (
	"flag"
	"time"
)

var (
	is_server *bool
	is_user   *bool
	listen    *string
)

func init() {
	is_server = flag.Bool("s", false, "server")
	is_user = flag.Bool("u", false, "getdata")
	listen = flag.String("l", ":12345", "listen addr")

	flag.Parse()
}

func main() {
	if *is_server {
		Service()
		return
	}

	if *is_user {
		User()
		return
	}

	// 客户端
	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()
	for {
		Client()
		<-t.C
	}

}
