package server

import (
	"errors"
	"fmt"
	"hostListen/base"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

func (l *Server) GetData(h int, result *map[string][]base.HostInfo) error {
	*result = base.HostData
	return nil
}
func liten() {
	t := time.NewTicker(time.Minute)
	for ; ; base.HostDataLock.Unlock() {
		<-t.C
		base.HostDataLock.Lock()
		if len(base.HostData) == 0 {
			continue
		}
		for _, v := range base.HostData {
			// last push time
			h := v[len(v)-1]
			t := int(time.Now().Unix()) - h.Date
			if t > 60 {
				// alert
				base.Mail.Set(base.UserMail, "host lost "+h.Sid, h.String()).Send()
				delete(base.HostData, h.Sid)
			}
		}
	}
}
func (l *Server) Save(h *base.HostInfo, result *string) error {
	*result = "I see"
	log.Println("recive a msg")
	if h.Sid == "" {
		return errors.New("sid is null")
	}

	base.HostDataLock.Lock()
	defer base.HostDataLock.Unlock()
	// find new host
	if base.HostData[h.Sid] == nil {
		base.HostData[h.Sid] = make([]base.HostInfo, 0)
		log.Println("find a new host")
		base.Mail.Set(base.UserMail, "HostListen find a new host", h.String()).Send()
	}

	base.HostData[h.Sid] = append(base.HostData[h.Sid], *h)
	if len(base.HostData[h.Sid]) > 90 {
		*result = "is much "
		// 转储
	}
	return nil
}
func init() {
	// 开启监听，失联报警
	go liten()
}
func Service() {
	//注册服务
	rpc.Register(new(Server))
	//绑定http协议
	rpc.HandleHTTP()
	// rpc.Accept()
	// rpc.DefaultServer.ServeHTTP()
	//监听服务
	fmt.Println("开始监听", *base.Listen)
	err := http.ListenAndServe(*base.Listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}
