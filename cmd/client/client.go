package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"conf-server/drpc/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const (
	DRPC_METHOD_CREATE_SCHEMA = 201

)

var (
	addr = flag.String("s", "127.0.0.1:8181", "defaule address")
	op = flag.String("t", "create_schema", "default api request")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDrpcServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var request *pb.Request
	var response *pb.Response
	switch(*op) {
	case "create_schema":
		createRequest := &pb.CreateSchemaReq{
			Name: fmt.Sprintf("schema-%d",rand.Int31()),
		}
		body,_:= json.Marshal(createRequest)
		request = &pb.Request{
		   Method: DRPC_METHOD_CREATE_SCHEMA,
		   Body:body,
		}
		response, err = c.DrpcFunc(ctx, request)
		if err != nil {
			log.Fatalf("could call DrpcFunc: %v", err)
		}
		createResp := &pb.CreateSchemaReq{}
		json.Unmarshal(response.Body,createResp)
		log.Info("createResponse=",createResp)
		break
	}


}
