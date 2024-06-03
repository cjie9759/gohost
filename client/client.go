package client

import (
	"crypto/tls"
	"crypto/x509"
	"gohost/base"
	hostinfo "gohost/hostInfo"
	"log"
	"net/rpc"
	"sync"
	"time"
)

func Con(lis string) (*tls.Conn, error) {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(base.SCert)

	cert, _ := tls.X509KeyPair(base.CCert, base.CKey)
	config := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
	}
	conn, err := tls.Dial("tcp", lis, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func Client() {
	if base.Name != "" {
		base.Name = base.Name + "----"
	}
	wg := &sync.WaitGroup{}
	for _, v := range base.Listen {
		wg.Add(1)
		go func(v string) {
			client(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
func client(Lis string) {
	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()
	var (
		conn   *tls.Conn
		err    error
		client *rpc.Client
	)
	c := func() {
		for {
			conn, err = Con(Lis)
			if err != nil {
				<-t.C
				log.Println("con err", err, "\n正在重新连接")
				continue
			}
			log.Println("con success")
			client = rpc.NewClient(conn)
			return
		}
	}
	f := func() {
		result := ""
		//调用方法
		err = client.Call("Server.Save", hostinfo.GetHostInfo(), &result)
		if err != nil {
			log.Println(err)
			c()
			return
		}
		log.Println("server return", result)
	}
	c()
	for {
		<-t.C
		go f()
	}
}
