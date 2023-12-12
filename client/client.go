package client

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"hostListen/base"
	"log"
	"net"
	"net/rpc"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
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
		err = client.Call("Server.Save", getHostInfo(), &result)
		if err != nil {
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

func getHostInfo() *base.HostInfo {
	cc, _ := cpu.Counts(false)
	ct, _ := cpu.Percent(time.Microsecond*3, false)
	cn, _ := cpu.Info()
	c := fmt.Sprint(cc, ct, cn[0].ModelName)

	d1, _ := disk.Usage("./")
	d := fmt.Sprint(d1.Used/1024/1024/1024, "G/", d1.Total/1024/1024/1024, "G/", d1.UsedPercent)

	meminfo, _ := mem.VirtualMemory()
	m := fmt.Sprint(meminfo.Total/1024/1024, "M/",
		meminfo.Used/1024/1024, "M/", meminfo.UsedPercent)

	hostinfo, _ := host.Info()
	hostname := base.Name + hostinfo.Hostname

	// netinfo := net.Addr.String()
	conn, _ := net.Dial("udp", "baidu.com:80")
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]

	sid1 := md5.Sum([]byte(
		fmt.Sprint(hostname, ip, "cj", cn[0].ModelName)))
	// sid := make([]byte, 16)

	data := &base.HostInfo{
		Sid:      hex.EncodeToString(sid1[:]),
		HostName: hostname,
		SysInfo:  fmt.Sprint(hostinfo.OS, "/", hostinfo.PlatformVersion),
		Ip:       ip,
		Mem:      m,
		Cpu:      c,
		Disk:     d,
		Date:     int(time.Now().Unix()),
	}
	return data
}
