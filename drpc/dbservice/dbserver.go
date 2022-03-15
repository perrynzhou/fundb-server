package dbservice

import (
	"fundb-server/drpc/pb"
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

const (
	drpcSockName = "drpc"
)

type DbService struct {
	pb.UnimplementedDrpcServiceServer
	drpcSockets []string
	count       uint64
}

func NewDbService(threadNum int) *DbService {
	dbService := &DbService{
		drpcSockets: make([]string, 0),
	}
	for i := 0; i < threadNum; i++ {
		drpcSocket := fmt.Sprintf("/tmp/%s_%d.sock", drpcSockName, i)
		dbService.drpcSockets = append(dbService.drpcSockets, drpcSocket)
		log.Info("run sock:", drpcSocket)
	}
	return dbService
}

func (cs *DbService) DrpcFunc(ctx context.Context, Req *pb.Request) (*pb.Response, error) {

	index := int(cs.count % uint64(len(cs.drpcSockets)))
	atomic.AddUint64(&cs.count, 1)
	c := NewClientConnection(cs.drpcSockets[index])
	if c == nil {
		return &pb.Response{}, errors.New("init connection failed")
	}
	if err := c.Connect(); err != nil {
		return &pb.Response{}, err

	}
	defer c.Close()
	log.Info("********recv request*******")
	log.Info("request method=", Req.Method, ",socket=", cs.drpcSockets[index])
	response, err := c.IssueCall(Req)

	if err != nil {
		response = &pb.Response{
			Body: []byte(err.Error()),
		}
	}
	return response, err
}
