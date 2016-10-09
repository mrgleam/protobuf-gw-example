package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	examples "github.com/mrgleam/protobuf-gw-example/pb"
	"google.golang.org/grpc"
)

func Run() error {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	examples.RegisterYourServiceServer(s, newEchoServer())

	s.Serve(l)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
