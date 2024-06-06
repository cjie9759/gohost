package main

import (
	"gohost/base"
	"gohost/client"
	"gohost/rpc"
	"testing"
)

func init() {
	base.IsTest = true
	base.Init()
}
func Test(t *testing.T) {
	base.Listen = []string{"127.0.0.1:12398"}
	rpc.TlsService()
}

func TestClient(t *testing.T) {
	client.NewClien(rpc.NewClient("127.0.0.1:12398")).Run()
}
