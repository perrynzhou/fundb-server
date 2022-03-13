package service

import (
	"conf-server/drpc/pb"
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	drpcSockName = "drpc"
)

type ConfService struct {
	pb.UnimplementedDrpcServiceServer
	drpcSocket string
}

func NewConfService() *ConfService {
	confService := &ConfService{}
	confService.drpcSocket = fmt.Sprintf("/tmp/%s.sock", drpcSockName)
	return confService
}

func (cs *ConfService) DrpcFunc(ctx context.Context, Req *pb.Request) (*pb.Response, error) {

	c := NewClientConnection(cs.drpcSocket)
	log.Info("********recv request*******")
	log.Info("request method=", Req.Method, ",socket=", cs.drpcSocket)
	response := &pb.Response{}
	if err := c.Connect(); err != nil {
		response.Body = []byte(err.Error())
		log.Info("err:", err)
		return response, err
	}
	defer c.Close()
	response, err := c.IssueCall(Req)
	if err != nil {
		response = &pb.Response{
			Body: []byte(err.Error()),
		}
	}
	return response, err

}
