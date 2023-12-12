package main

import (
	"hostListen/base"
	"hostListen/client"
	"hostListen/server"
	"testing"
)

func init() {
	base.IsTest = true
	base.Init()
}
func Test(t *testing.T) {
	base.Listen = []string{"127.0.0.1:12398"}
	go server.Listen()
	server.TlsService()
}

func TestClient(t *testing.T) {
	base.Listen = []string{"127.0.0.1:12398"}
	client.Client()
}
