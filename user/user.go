package user

import (
	"fmt"
	"hostListen/base"
	"log"
	"net/rpc"
)

func showHostData() {
	for k, v := range base.HostData {
		fmt.Println("Host", k)
		for _, v1 := range v {
			fmt.Println()
			fmt.Println(v1.String())
		}
	}
}
func User() {
	//连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", *base.Listen)
	if err != nil {
		log.Println(err)
	}

	//调用方法
	result := base.HostData
	err = conn.Call("Server.GetData", 1, &result)
	showHostData()

	if err != nil {
		log.Println(err)
		return
	}
}
