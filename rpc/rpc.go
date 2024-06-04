package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gohost/base"
	hostinfo "gohost/hostInfo"
	"gohost/server"
	"log"
	"net/rpc"
	sync "sync"
)

func TlsService() {
	se := server.Server{}
	se.Run()

	//注册服务
	s := rpc.NewServer()
	s.Register(se)

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

//	type ServiceServer interface {
//		Save(context.Context, *HostInfo) (*HostInfoRes, error)
//		Get(context.Context, *HostInfoRes) (*HostInfo, error)
//	}
type Server struct {
	s *server.Server
}

func (s *Server) Save(h *hostinfo.HostInfo, r *res) error {
	return s.s.Save(h)
}
func (s *Server) GetData(r *req, res map[string]*hostinfo.HostInfo) error {
	res = s.s.GetData()
	return nil
}

type res string
type req string
