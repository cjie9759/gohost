package main

import (
	"gohost/base"
	"gohost/client"
	"gohost/rpc"
	"gohost/user"
	"sync"
)

func main() {
	base.Init()

	switch {
	case base.Is_server:
		rpc.TlsService()
	case base.Is_user:
		user.User()
	default:
		wg := &sync.WaitGroup{}
		wg.Add(len(base.Listen))
		for _, v := range base.Listen {
			s := v
			go func() {
				defer wg.Done()
				client.NewClien(rpc.NewClient(s)).Run()
			}()
		}
		wg.Wait()
	}

	// 客户端
	// <-t.C
	// t := time.NewTicker(time.Minute / 10)
	// defer t.Stop()
	// for {
	// }

}
