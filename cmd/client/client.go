package main

import (
	"conf-server/drpc/pb"
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

const (
	DRPC_METHOD_CREATE_SCHEMA = 201
)

var (
	addr  = flag.String("s", "127.0.0.1:50051", "defaule address")
	op    = flag.String("t", "create_schema", "default api request")
	count = flag.Int("n", 1, "default run times")
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()
	var request *pb.Request
	var response *pb.Response
	rand.Seed(time.Now().Unix())
	r := rand.Int63() % 1024
	for i := 0; i < *count; i++ {
		switch *op {
		case "create_schema":
			createRequest := &pb.CreateSchemaReq{
				Name: fmt.Sprintf("schema-%d", r+int64(rand.Int())),
			}
			// createRequest
			body, _ := proto.Marshal(createRequest)
			request = &pb.Request{
				Method: DRPC_METHOD_CREATE_SCHEMA,
				Body:   body,
			}
			response, err = c.DrpcFunc(ctx, request)
			if err != nil {
				log.Fatalf("could call DrpcFunc: %v", err)
			}
			createResp := &pb.CreateSchemaReq{}
			proto.Unmarshal(response.Body, createResp)
			log.Info("createResponse:", createResp)
		}
	}
}
