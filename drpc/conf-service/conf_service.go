package conf_service

import (
	"conf-server/drpc/pb"
)

const (
	DRPC_METHOD_CREATE_SCHEMA = 101
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
