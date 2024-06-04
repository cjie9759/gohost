package rpc

import (
	context "context"
	"gohost/base"
	hostinfo "gohost/hostInfo"
	"gohost/server"
	"log"
	"net"
	sync "sync"

	"github.com/shirou/gopsutil/mem"
	grpc "google.golang.org/grpc"
)

func RPCService() {

	se := server.Server{}
	se.Run()

	wg := &sync.WaitGroup{}
	for _, v := range base.Listen {
		wg.Add(1)
		go func(v string) {
			lis, err := net.Listen("tcp", v)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			gs := grpc.NewServer()
			RegisterServiceServer(gs, &Server{s: &se})
			if err := gs.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
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

func (s *Server) Save(c context.Context, h *HostInfo) (*HostInfoRes, error) {

	s.s.Save(&hostinfo.HostInfo{
		Sid:      h.Sid,
		HostName: h.HostName,
		SysInfo:  h.SysInfo,
		Ip:       h.Ip,
		Sip:      h.Sip,
		Mem:      h.Mem.(*mem.VirtualMemoryStat),
		Host:     h.Host,
		Cpu:      h.Cpu,
		Disk:     h.Disk,
		Date:     h.Date,
		Time:     h.Time,
		LTime:    h.LTime,
	})

}
func (s *Server) Get(c context.Context, h *HostInfoRes) (*HostInfo, error) {}
