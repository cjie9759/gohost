package base

import "flag"

type HostInfo struct {
	Sid      string
	HostName string
	SysInfo  string
	Ip       string
	Mem      string
	Cpu      string
	Disk     string
	Date     int
}

var HostData = make(map[string][]HostInfo)

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
