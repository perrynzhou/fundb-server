package main

import (
	"conf-server/drpc/conf-service"
	"conf-server/drpc/pb"
	"conf-server/util"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	sockFile = flag.String("s", "/tmp/drpc_socket.sock", "local socekt file")
	port     = flag.Int("p", 50051, "The server port")
)

func init() {
	logHook := util.NewHook()
	logHook.Field = "line"
	log.AddHook(logHook)
}

func main() {
	flag.Parse()
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	confService := conf_service.NewConfService(*sockFile)
	pb.RegisterDrpcServiceServer(s, confService)
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	for {
		select {
		case <-sigs:
			log.Info("stop meta server....")
			return
		}
	}
}
