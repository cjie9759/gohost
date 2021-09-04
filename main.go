package main

import (
	"flag"
	"time"
)

var (
	is_server *bool
	listen    *string
)

func init() {
	is_server = flag.Bool("s", false, "server")
	listen = flag.String("l", ":12345", "listen addr")

	flag.Parse()
}

func main() {
	if *is_server {
		Service()
		return
	}

	// 客户端
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for {
		Client()
		<-t.C
	}

}
