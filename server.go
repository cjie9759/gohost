package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type HostInfo struct {
	Sid      string
	HostName string
	SysInfo  string
	Ip       string
	Mem      string
	Cpu      string
	Disk     string
	Date     int
}
type Server struct {
}

//函数名首字母大写
//第一个参数为接收参数，第二个参数是返回结果，必须是指针类型
//函数结果必须返回一个error
func (l *Server) Save(h *HostInfo, result *string) error {
	*result = "I receive"
	fmt.Println()
	fmt.Println("Sid", h.Sid)
	fmt.Println("HostName", h.HostName)
	fmt.Println("SysInfo", h.SysInfo)
	fmt.Println("Ip", h.Ip)
	fmt.Println("Mem", h.Mem)
	fmt.Println("Cpu", h.Cpu)
	fmt.Println("Disk", h.Disk)
	fmt.Println("Date", h.Date)
	return nil
}

func Service() {
	//注册服务
	rpc.Register(new(Server))
	//绑定http协议
	rpc.HandleHTTP()
	//监听服务
	fmt.Println("开始监听", *listen)
	err := http.ListenAndServe(*listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
