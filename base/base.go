package base

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"sync"
)

type HostInfo struct {
	Sid      string
	HostName string
	SysInfo  string
	Ip       string
	Sip      string
	Mem      string
	Cpu      string
	Disk     string
	Date     int
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
	fmt.Fprintln(a, "Date", t.Date)
	return a.Bytes()
}
func (t *HostInfo) String() string {
	return string(t.Bytes())
}

var HostData = make(map[string][]HostInfo)

var HostDataLock = new(sync.RWMutex)

// var HostData = make(syncmap[string][]HostInfo)

var (
	Is_server *bool
	Is_user   *bool
	Listen    *string
)

func init() {
	Is_server = flag.Bool("s", false, "server")
	Is_user = flag.Bool("u", false, "getdata")
	Listen = flag.String("l", ":12345", "listen addr")

	flag.Parse()
}

// //go:embed pem/fullchain.pem
// var Cert []byte

// //go:embed pem/privkey.pem
// var Key []byte

//go:embed pem/client.crt
var CCert []byte

//go:embed pem/client.key
var CKey []byte

//go:embed pem/server.crt
var SCert []byte

//go:embed pem/server.key
var SKey []byte
