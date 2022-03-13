package service

import (
	"conf-server/drpc/pb"
	"context"
	"fmt"
	"math"
	"sync"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

const (
	drpcSockName = "drpc"
	syncSockName = "sync"
)

type ConfService struct {
	pb.UnimplementedDrpcServiceServer
	drpcSockets []string
	syncSockets []string
	counter     uint64
	mutex       *sync.Mutex
}

func NewConfService(threadNum int) *ConfService {
	confService := &ConfService{}
	for i := 0; i < threadNum; i++ {
		if confService.drpcSockets == nil {
			confService.drpcSockets = make([]string, 0)
		}
		if confService.syncSockets == nil {
			confService.syncSockets = make([]string, 0)
		}
		drpcSocket := fmt.Sprintf("/tmp/%s_%d.sock", drpcSockName, i)
		syncSockets := fmt.Sprintf("/tmp/%s_%d.sock", syncSockName, i)

		confService.drpcSockets = append(confService.drpcSockets, drpcSocket)
		confService.syncSockets = append(confService.drpcSockets, syncSockets)

	}
	confService.mutex = &sync.Mutex{}
	return confService
}

func (cs *ConfService) fetchNextSocket(socketFiles []string) string {
	if cs.counter >= math.MaxUint64-1 {
		cs.mutex.Lock()
		cs.counter = 0
		cs.mutex.Unlock()
	}
	atomic.AddUint64(&cs.counter, 1)
	index := int(cs.counter % uint64(len(socketFiles)))
	log.Info("Round-Robin socket:",socketFiles[index])
	return socketFiles[index]

}
func (cs *ConfService) DrpcFunc(ctx context.Context, Req *pb.Request) (*pb.Response, error) {
	log.Info("********recv request*******")
	log.Info("request method=", Req.Method)
	socketFile := cs.fetchNextSocket(cs.drpcSockets)
	c := NewClientConnection(socketFile)
	c.Connect()
	response, err := c.IssueCall(Req)
	if err != nil {
		response = &pb.Response{
			Body: []byte(err.Error()),
		}
	}
	return response, err

}
