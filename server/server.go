package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"hostListen/base"
	"log"
	"net/rpc"
	"sync"
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
			if t > int(base.LosTime.Nanoseconds()) {
				// alert
				go base.Notifys.Send("host lost " + h.HostName + "  " + h.Sid + h.String())
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
		go base.Notifys.Send("host find " + h.HostName + "  " + h.Sid + h.String())

	}
	// 使用系统时间
	h.Date = int(time.Now().Unix())
	base.HostData[h.Sid] = append(base.HostData[h.Sid], *h)
	if len(base.HostData[h.Sid]) > 90 {
		*result = "is much "
		base.HostData[h.Sid] = base.HostData[h.Sid][80:]
		// 转储
	}
	return nil
}

func Init() {
	// 开启监听，失联报警
	go liten()
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

	// https
	// hs := &http.Server{
	// 	Addr:           *base.Listen,
	// 	Handler:        s,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// 	TLSConfig:      config,
	// }
	// fmt.Println("开始监听", *base.Listen)
	// err := hs.ListenAndServeTLS("", "")

}

// func Service() {
// 	//注册服务
// 	s := rpc.NewServer()
// 	s.Register(new(Server))
// 	hs := &http.Server{
// 		Addr:           base.Listen,
// 		Handler:        s,
// 		ReadTimeout:    10 * time.Second,
// 		WriteTimeout:   10 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	fmt.Println("开始监听", base.Listen)
// 	err := hs.ListenAndServe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
