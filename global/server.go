package global

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GRPCConfig struct {
	Name         string
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func RunGrpcServer(c *GRPCConfig) error {
	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		logrus.Fatal("cannot listen", c.Name, err)
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	logrus.Info("SERVER STARTED", c.Name, "  ON  ", c.Addr)
	return s.Serve(listen)
}
