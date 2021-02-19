package main

import (
	"net"
	"os"

	protos "github.com/akashkumar8/micproject/protos"
	"github.com/akashkumar8/micproject/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := hclog.Default()

	gs := grpc.NewServer()

	cs := server.NewCourse(log)

	protos.RegisterCourseServer(gs, cs)
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
