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

var sm = &sync.Map{}

type Server struct {
}

func (l *Server) GetData(h int, result *[]base.HostInfo) error {
	*result = make([]base.HostInfo, 10)
	return nil
}
func (l *Server) Save(h *base.HostInfo, result *string) error {
	*result = "I see"
	log.Println("recive a msg")
	if h.Sid == "" {
		return errors.New("sid is null")
	}

	// find new host
	_, ok := sm.Load(h.Sid)
	if !ok {
		log.Println("find a new host")
		base.Notifys.Send("host find " + h.HostName + "  " + h.Sid + h.String())
	}

	// 使用系统时间
	h.Date = int(time.Now().Unix())
	sm.Store(h.Sid, h)

	return base.DB.Save(h).Error
}
func Listen() {
	t := time.NewTicker(time.Minute)
	for {
		<-t.C

		sm.Range(func(key, value any) bool {
			h := value.(*base.HostInfo)
			t := int(time.Now().Unix()) - h.Date
			// fmt.Println(t, int(base.LosTime.Seconds()), h.Date)
			if t > int(base.LosTime.Seconds()) {
				// alert
				base.Notifys.Send("host lost " + h.HostName + "  " + h.Sid + h.String())
				sm.Delete(key)
			}
			return true
		})
	}
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
