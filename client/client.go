package client

import (
	"crypto/md5"
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

func Client() {
	//连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", *base.Listen)
	if err != nil {
		log.Println(err)
	}

	//调用方法
	result := ""
	err = conn.Call("Server.Save", getHostInfo(), &result)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("server return", result)
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
