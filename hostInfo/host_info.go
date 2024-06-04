package hostinfo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gohost/base"
	"net"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gorm.io/gorm"
)

type HostInfo struct {
	gorm.Model
	Sid      string
	HostName string
	SysInfo  string
	Ip       string
	Sip      string
	Mem      *mem.VirtualMemoryStat `gorm:"type:josnb;serializer:json"`
	Host     *host.InfoStat         `gorm:"type:josnb;serializer:json"`
	Cpu      CPUinfo                `gorm:"type:josnb;serializer:json"`
	Disk     *disk.UsageStat        `gorm:"type:josnb;serializer:json"`
	Date     int                    `gorm:"index"`
	Time     time.Time
	LTime    time.Time
}
type CPUinfo struct {
	Count   int
	Percent float64
	Info    cpu.InfoStat
}

func (t *HostInfo) Bytes() []byte {
	a := bytes.NewBuffer(nil)
	fmt.Fprintln(a, "Sid", t.Sid)
	fmt.Fprintln(a, "HostName", t.HostName)
	fmt.Fprintln(a, "SysInfo", t.SysInfo)
	fmt.Fprintln(a, "Ip", t.Ip)
	fmt.Fprintln(a, "Mem", t.Mem)
	fmt.Fprintln(a, "Cpu", t.Cpu)
	fmt.Fprintln(a, "Disk", t.Disk)
	d := time.Unix(int64(t.Date), 0).Local().Format("01/02 15:04:05")
	fmt.Fprintln(a, "Date", d)
	return a.Bytes()
}
func (t *HostInfo) String() string {
	return string(t.Bytes())
}

func GetHostInfo() *HostInfo {
	cc, _ := cpu.Counts(false)
	ct, _ := cpu.Percent(time.Microsecond*100, false)
	cn, _ := cpu.Info()

	d, _ := disk.Usage("./")

	meminfo, _ := mem.VirtualMemory()

	hostinfo, _ := host.Info()
	hostname := base.Name + hostinfo.Hostname

	// netinfo := net.Addr.String()
	conn, _ := net.Dial("udp", "jd.com:80")
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]

	sid := md5.Sum([]byte(
		fmt.Sprint(hostname, ip, "cj", cn[0].ModelName)))

	data := &HostInfo{
		Sid:      hex.EncodeToString(sid[:]),
		HostName: hostname,
		SysInfo:  fmt.Sprint(hostinfo.OS, "/", hostinfo.PlatformVersion),
		Ip:       ip,
		Mem:      meminfo,
		Host:     hostinfo,
		Cpu: CPUinfo{
			Count:   cc,
			Percent: ct[0],
			Info:    cn[0]},
		Disk:  d,
		Date:  int(time.Now().Unix()),
		Time:  time.Now(),
		LTime: time.Now(),
	}
	return data
}
