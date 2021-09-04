package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

//net/rpc库使用encoding/gob进行编解码，只支持golang调用

type Params struct {
	Name string
}

type Love struct {
}

//函数名首字母大写
//第一个参数为接收参数，第二个参数是返回结果，必须是指针类型
//函数结果必须返回一个error
func (l *Love) Confession(p Params, result *string) error {
	*result = "I love you, " + p.Name
	return nil
}

func main() {
	//注册服务
	love := new(Love)
	rpc.Register(love)
	//绑定http协议
	rpc.HandleHTTP()
	//监听服务
	fmt.Println("开始监听8888端口...")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
