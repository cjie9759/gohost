package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"gohost/base"
	hostinfo "gohost/hostInfo"
	"log"
	"net/rpc"
	"sync"
	"time"
)

// var sm = &sync.Map{}

type Server struct {
	data map[string]*hostinfo.HostInfo
	ch   chan func()
}

func (l *Server) GetData() map[string]*hostinfo.HostInfo {
	return l.data
}

func (l *Server) Run(h *hostinfo.HostInfo) error {

	if l.data == nil {
		l.data = make(map[string]*hostinfo.HostInfo, 10)
	}

	if l.ch == nil {
		l.ch = make(chan func(), 1000)
	}

	go func() {
		f := <-l.ch
		f()
	}()

	// 离线监测
	go func() {
		t := time.NewTicker(time.Minute)
		for {
			<-t.C

			l.ch <- func() {
				deleteKey := []string{}
				for k, h := range l.data {
					t := int(time.Now().Unix()) - h.Date
					if t > int(base.LosTime.Seconds()) {
						// alert
						base.Notifys.Send("host lost " + h.HostName + "  " + h.Sid + h.String())
						deleteKey = append(deleteKey, k)
					}
				}

				for _, v := range deleteKey {
					delete(l.data, v)
				}
			}
		}
	}()

	return nil
}

func (l *Server) Save(h *hostinfo.HostInfo) error {
	if h.Sid == "" {
		return errors.New("sid is null")
	}
	revTime := time.Now()
	h.Date = int(revTime.Unix())
	l.ch <- func() {

		// find new host
		if l.data[h.Sid] == nil && time.Since(base.Uptime) > time.Minute {
			log.Println("find a new host")
			base.Notifys.Send("host find " + h.HostName + "  " + h.Sid + h.String())
		}

		l.data[h.Sid] = h
		log.Printf("Save host %s in %v", h.Sid, time.Since(revTime))
	}

	return base.DB.Save(h).Error
}

func TlsService() {
	//注册服务
	s := rpc.NewServer()
	s.Register(new(Server))

	cert, _ := tls.X509KeyPair(base.SCert, base.SKey)
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(base.CCert)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	wg := &sync.WaitGroup{}
	for _, v := range base.Listen {
		wg.Add(1)
		go func(v string) {
			l, err := tls.Listen("tcp", v, config)
			fmt.Println("开始监听", v)
			s.Accept(l)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()

}
