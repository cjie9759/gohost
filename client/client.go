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
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func Con() (*tls.Conn, error) {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(base.SCert)

	cert, _ := tls.X509KeyPair(base.CCert, base.CKey)
	config := &tls.Config{
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
	}
	conn, err := tls.Dial("tcp", *base.Listen, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func Client() {
	t := time.NewTicker(time.Minute / 10)
	defer t.Stop()
	f := func() {
		conn, err := Con()
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		client := rpc.NewClient(conn)

		//调用方法
		result := ""
		err = client.Call("Server.Save", getHostInfo(), &result)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("server return", result)
	}
	for {
		<-t.C
		f()
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
	hostname := hostinfo.Hostname

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
