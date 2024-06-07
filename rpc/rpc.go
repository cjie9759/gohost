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
	"sync"
	"time"
)

func TlsService() {
	se := server.Server{}
	se.Run()

	//注册服务
	s := rpc.NewServer()
	s.Register(&Server{&se})

	cert, _ := tls.X509KeyPair(base.SCert, base.SKey)
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(base.CCert)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	t := time.NewTicker(time.Millisecond)
	wg := &sync.WaitGroup{}
	wg.Add(len(base.Listen))
	for _, v := range base.Listen {
		lis := v
		go func() {
			defer wg.Done()
			for {
				l, err := tls.Listen("tcp", lis, config)
				fmt.Println("开始监听", lis)
				s.Accept(l)
				if err != nil {
					log.Fatalln(err)
				}
				<-t.C
			}
		}()
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

func (s *Server) Save(h *hostinfo.HostInfo, r *Res) error {
	return s.s.Save(h)
}
func (s *Server) GetData(r *Req, res map[string]*hostinfo.HostInfo) error {
	res = s.s.GetData()
	return nil
}

type Res string
type Req string

func NewClient(lis string) server.ServerInterface {
	return &client{lis: lis}
}

type client struct {
	cfg *tls.Config
	c   *rpc.Client
	lis string
}

func (c *client) con() {

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(base.SCert)

	cert, _ := tls.X509KeyPair(base.CCert, base.CKey)
	c.cfg = &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
	}

	var (
		conn *tls.Conn
		err  error
	)
	t := time.NewTicker(time.Millisecond)
	for {
		conn, err = tls.Dial("tcp", c.lis, c.cfg)
		if err != nil {
			log.Println(err)
			<-t.C
			continue
		}
		if conn != nil {
			break
		}
	}
	c.c = rpc.NewClient(conn)
}
func (c *client) Save(h *hostinfo.HostInfo) error {
	if c.c == nil {
		c.con()
	}

	res := ""
	err := c.c.Call("Server.Save", h, &res)
	if err != nil {
		log.Println(err)
		c.con()
		return err
	}
	log.Println("server return", res)
	return nil
}
func (c *client) GetData() map[string]*hostinfo.HostInfo {
	if c.c == nil {
		c.con()
	}

	res := map[string]*hostinfo.HostInfo{}
	err := c.c.Call("Server.GetData", "", res)
	if err != nil {
		log.Println(err)
		c.con()
		return nil
	}
	log.Println("server return", res)
	return nil
}
