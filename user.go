package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func showHostData() {
	for k, v := range hostData {
		fmt.Println("Host", k)
		for _, v1 := range v {
			fmt.Println()
			fmt.Println("Sid", v1.Sid)
			fmt.Println("HostName", v1.HostName)
			fmt.Println("SysInfo", v1.SysInfo)
			fmt.Println("Ip", v1.Ip)
			fmt.Println("Mem", v1.Mem)
			fmt.Println("Cpu", v1.Cpu)
			fmt.Println("Disk", v1.Disk)
			fmt.Println("Date", v1.Date)
		}
	}
}
func User() {
	//连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", *listen)
	if err != nil {
		log.Println(err)
	}

	//调用方法
	result := hostData
	err = conn.Call("Server.GetData", 1, &result)
	showHostData()

	if err != nil {
		log.Println(err)
		return
	}
}
