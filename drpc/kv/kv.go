package kv

import (
	"conf-server/drpc/pb"
	"context"
)

const (
	DRPC_METHOD_CREATE_SCHEMA = 101
)

type KvService struct {
	pb.UnimplementedKvServiceServer
	localSockFile string
}

func NewKvService(sockFile string) *KvService {
	return &KvService{
		localSockFile: sockFile,
	}
}

/*

rpc CreateSchema(CreateSchemaReq) returns(CreateSchemaResp) {}
	rpc QuerySchema(QuerySchemaReq) returns(QuerySchemaResp) {}
	rpc DeleteSchema(DeleteSchemaReq) returns(DeleteSchemaResp) {}
	rpc PutConf(PutConfReq) returns(PutConfResp) {}
	rpc GetConf(GetConfReq) returns(GetConfResp) {}
	rpc DeleteConf(DeleteConfReq) returns(DeleteConfResp) {}
*/
func (s *KvService) CreateSchema(c context.Context, req *pb.CreateSchemaReq) (*pb.CreateSchemaResp, error) {
	return nil, nil
}

func (s *KvService) QuerySchema(c context.Context, req *pb.QuerySchemaReq) (*pb.QuerySchemaResp, error) {
	return nil, nil
}

func (s *KvService) DeleteSchema(c context.Context, req *pb.DeleteSchemaReq) (*pb.DeleteSchemaResp, error) {
	return nil, nil
}

func (s *KvService) PutConf(c context.Context, req *pb.PutConfReq) (*pb.PutConfResp, error) {
	return nil, nil
}

func (s *KvService) GetConf(c context.Context, req *pb.GetConfReq) (*pb.GetConfResp, error) {
	return nil, nil
}

func (s *KvService) DeleteConf(c context.Context, req *pb.DeleteConfReq) (*pb.DeleteConfResp, error) {
	return nil, nil
}
