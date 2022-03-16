package main

import (
	"context"
	"encoding/json"
	_ "encoding/json"
	"flag"
	"fmt"
	"fundb-server/drpc/pb"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

const (
	DRPC_METHOD_CREATE_SCHEMA = 201
	DRPC_METHOD_DROP_SCHEMA   = 202

	DRPC_METHOD_PUT_KV = 301
	DRPC_METHOD_GET_KV = 302
	DRPC_METHOD_DEL_KV = 303
)

var (
	addr  = flag.String("s", "127.0.0.1:50051", "defaule address")
	op    = flag.String("t", "create_schema", "default api request")
	count = flag.Int("n", 1, "default run times")
)

type TData struct {
	Id   int    `json:"id"`
	Port int    `json:"port"`
	Path string `json:"path"`
}

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
	for i := 0; i < *count; i++ {
		switch *op {
		case "put_kv":
			val := &TData{
				Id:   rand.Int(),
				Port: rand.Intn(*count) + 4000,
				Path: fmt.Sprintf("/tmp/data-%d", rand.Int()),
			}
			b, err := json.Marshal(val)
			if err != nil {
				log.Error(err)
				break
			}
			putKvRequest := &pb.PutKvReq{
				SchemaName: fmt.Sprintf("schema-%d", i),
				Key:        fmt.Sprintf("key-%d", rand.Int31()),
				Value:      string(b),
			}
			// createRequest
			body, _ := proto.Marshal(putKvRequest)
			request = &pb.Request{
				Method: DRPC_METHOD_PUT_KV,
				Body:   body,
			}
			response, err = c.DrpcFunc(ctx, request)
			if err != nil {
				log.Fatalf("could call DrpcFunc: %v", err)
			}
			putKvResp := &pb.PutKvResp{}
			proto.Unmarshal(response.Body, putKvResp)
			log.Info(putKvResp)
		case "drop_schema":
			dropRequest := &pb.DropSchemaReq{
				Name: fmt.Sprintf("schema-%d", i),
			}
			// createRequest
			body, _ := proto.Marshal(dropRequest)
			request = &pb.Request{
				Method: DRPC_METHOD_DROP_SCHEMA,
				Body:   body,
			}
			response, err = c.DrpcFunc(ctx, request)
			if err != nil {
				log.Fatalf("could call DrpcFunc: %v", err)
			}
			dropResp := &pb.DropSchemaResp{}
			proto.Unmarshal(response.Body, dropResp)
			log.Info(dropResp)
		case "create_schema":
			createRequest := &pb.CreateSchemaReq{
				Name: fmt.Sprintf("schema-%d", i),
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
			createResp := &pb.CreateSchemaResp{}
			proto.Unmarshal(response.Body, createResp)
			log.Info(createResp)
			/*
				b, err := json.MarshalIndent(createResp, " ", " ")
				if err != nil {
					log.Error(err)
				} else {
					log.Info("createResponse:", string(b))
				}
			*/
		}
	}
}
