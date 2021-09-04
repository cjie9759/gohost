package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Name string
}

func main() {
	//连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	//调用方法
	result := ""
	err = conn.Call("Love.Confession", Params{"BaoBao"}, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
