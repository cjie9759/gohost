package main

import (
	"gohost/base"
	"gohost/client"
	hostinfo "gohost/hostInfo"
	"gohost/rpc"
	"log"
	"testing"
)

func init() {
	base.IsTest = true
}
func Test(t *testing.T) {
	base.Is_server = true
	base.Init()
	err := base.DB.AutoMigrate(&hostinfo.HostInfo{})
	if err != nil {
		log.Panic("db connect fail:", err)
	}

	base.Listen = []string{"127.0.0.1:12398"}
	rpc.TlsService()
}

func TestClient(t *testing.T) {
	base.Init()
	client.NewClien(rpc.NewClient("127.0.0.1:12398")).Run()
}
