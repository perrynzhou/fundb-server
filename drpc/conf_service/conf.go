package conf_service

import (
	"conf-server/drpc/pb"
	"context"
	log "github.com/sirupsen/logrus"
)


type ConfService struct {
	pb.UnimplementedDrpcServiceServer
	localSockFile string
}

func NewConfService(sockFile string) *ConfService {
	return &ConfService{
		localSockFile: sockFile,
	}
}

func (cs *ConfService)DrpcFunc(ctx context.Context,Req *pb.Request)  (*pb.Response,error) {
	log.Info("********recv request*******")
	log.Info("request method=",Req.Method)
	c := NewClientConnection(cs.localSockFile)
	c.Connect()
	response,err :=c.IssueCall(Req)
	if err != nil {
		response = &pb.Response{
			Body:[]byte(err.Error()),
		}
	}
	return response,err

}


