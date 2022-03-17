package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"fundb-server/conf"
	"fundb-server/drpc/pb"
	"fundb-server/util"
	"math/rand"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

const (
	DRPC_METHOD_CREATE_SCHEMA = 201
	DRPC_METHOD_DROP_SCHEMA   = 202
	DRPC_METHOD_QUERY_SCHEMA  = 203

	DRPC_METHOD_PUT_KV = 301
	DRPC_METHOD_GET_KV = 302
	DRPC_METHOD_DEL_KV = 303
)

var (
	opTypes = []int{DRPC_METHOD_CREATE_SCHEMA, DRPC_METHOD_DROP_SCHEMA, DRPC_METHOD_PUT_KV, DRPC_METHOD_GET_KV, DRPC_METHOD_DEL_KV, DRPC_METHOD_QUERY_SCHEMA}
	opMap   = map[int]string{
		DRPC_METHOD_CREATE_SCHEMA: "create_schema",
		DRPC_METHOD_DROP_SCHEMA:   "drop_schema",
		DRPC_METHOD_QUERY_SCHEMA:  "query_schema",
		DRPC_METHOD_PUT_KV:        "put_kv",
		DRPC_METHOD_GET_KV:        "get_kv",
		DRPC_METHOD_DEL_KV:        "del_kv",
	}
)
var (
	addr     = flag.String("s", "127.0.0.1:50051", "defaule address")
	count    = flag.Int("n", 1000, "default run times")
	interval = flag.Int("i", 1, "default op interval millseconds")
)

type OpMetric struct {
	Succ  uint64
	Fail  uint64
	Total uint64
}
type TData struct {
	Id   int    `json:"id"`
	Port int    `json:"port"`
	Name string `json:"name"`
	Uid  string `json:"uid"`
}

func init() {
	logHook := util.NewHook()
	logHook.Field = "line"
	log.AddHook(logHook)
	conf.GenConfTemplate()
}
func querySchmea(c pb.DrpcServiceClient, ctx context.Context, n int32) int32 {
	i := rand.Int31n(n)
	queryRequest := &pb.QuerySchemaReq{
		Name: fmt.Sprintf("schema-%d", i),
	}
	// createRequest
	body, _ := proto.Marshal(queryRequest)
	request := &pb.Request{
		Method: DRPC_METHOD_QUERY_SCHEMA,
		Body:   body,
	}
	response, err := c.DrpcFunc(ctx, request)
	if err != nil {
		log.Fatalf("could call DrpcFunc: %v", err)
	}
	queryResp := &pb.QuerySchemaResp{}
	proto.Unmarshal(response.Body, queryResp)
	// log.Info(createResp)
	return queryResp.Code
}
func createSchmea(c pb.DrpcServiceClient, ctx context.Context, n int32) int32 {
	i := rand.Int31n(n)
	createRequest := &pb.CreateSchemaReq{
		Name: fmt.Sprintf("schema-%d", i),
	}
	// createRequest
	body, _ := proto.Marshal(createRequest)
	request := &pb.Request{
		Method: DRPC_METHOD_CREATE_SCHEMA,
		Body:   body,
	}
	response, err := c.DrpcFunc(ctx, request)
	if err != nil {
		log.Fatalf("could call DrpcFunc: %v", err)
	}
	createResp := &pb.CreateSchemaResp{}
	proto.Unmarshal(response.Body, createResp)
	// log.Info(createResp)
	return createResp.Code
}
func dropSchmea(c pb.DrpcServiceClient, ctx context.Context, n int32) int32 {
	i := rand.Int31n(n)
	dropRequest := &pb.DropSchemaReq{
		Name: fmt.Sprintf("schema-%d", i),
	}
	// createRequest
	body, _ := proto.Marshal(dropRequest)
	request := &pb.Request{
		Method: DRPC_METHOD_DROP_SCHEMA,
		Body:   body,
	}
	response, err := c.DrpcFunc(ctx, request)
	if err != nil {
		log.Fatalf("could call DrpcFunc: %v", err)
	}
	dropResp := &pb.DropSchemaResp{}
	proto.Unmarshal(response.Body, dropResp)
	//log.Info(dropResp)
	return dropResp.Code
}
func putKv(c pb.DrpcServiceClient, ctx context.Context, n int32) int32 {
	i := rand.Int31n(n)
	val := &TData{
		Id:   rand.Int(),
		Port: rand.Intn(*count) + 4000,
		Uid:  uuid.New().String(),
		Name: fmt.Sprintf("data-%d-perrynzhou", rand.Int()),
	}
	b, err := json.Marshal(val)
	if err != nil {
		log.Error(err)
		return -1
	}
	putKvRequest := &pb.PutKvReq{
		SchemaName: fmt.Sprintf("schema-%d", i),
		Key:        fmt.Sprintf("key-%d", i),
		Value:      string(b),
	}
	// createRequest
	body, _ := proto.Marshal(putKvRequest)
	request := &pb.Request{
		Method: DRPC_METHOD_PUT_KV,
		Body:   body,
	}
	response, err := c.DrpcFunc(ctx, request)
	if err != nil {
		log.Fatalf("could call DrpcFunc: %v", err)
	}
	putKvResp := &pb.PutKvResp{}
	proto.Unmarshal(response.Body, putKvResp)
	//log.Info(putKvResp)
	return putKvResp.Code
}
func getKv(c pb.DrpcServiceClient, ctx context.Context, n int32) int32 {
	i := rand.Int31n(n)

	getKvRequest := &pb.GetKvReq{
		SchemaName: fmt.Sprintf("schema-%d", i),
		Key:        fmt.Sprintf("key-%d", i),
	}
	// createRequest
	body, _ := proto.Marshal(getKvRequest)
	request := &pb.Request{
		Method: DRPC_METHOD_GET_KV,
		Body:   body,
	}
	response, err := c.DrpcFunc(ctx, request)
	if err != nil {
		log.Fatalf("could call DrpcFunc: %v", err)
	}
	getKvResp := &pb.GetKvResp{}
	proto.Unmarshal(response.Body, getKvResp)
	//log.Info(getKvResp)
	return getKvResp.Code
}
func delKv(c pb.DrpcServiceClient, ctx context.Context, n int32) int32 {
	i := rand.Int31n(n)
	delKvRequest := &pb.DelKvReq{
		SchemaName: fmt.Sprintf("schema-%d", i),
		Key:        fmt.Sprintf("key-%d", i),
	}
	// createRequest
	body, _ := proto.Marshal(delKvRequest)
	request := &pb.Request{
		Method: DRPC_METHOD_DEL_KV,
		Body:   body,
	}
	response, err := c.DrpcFunc(ctx, request)
	if err != nil {
		log.Fatalf("could call DrpcFunc: %v", err)
	}
	delKvResp := &pb.DelKvResp{}
	proto.Unmarshal(response.Body, delKvResp)
	return delKvResp.Code
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

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	n := int32(*count)
	metricMap := make(map[string]*OpMetric)
	for _, v := range opMap {
		metricMap[v] = &OpMetric{}
	}
	go func(mp map[string]*OpMetric) {
		ticker := time.NewTicker(time.Duration(1) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				for _, v := range opTypes {
					opName := opMap[v]
					metric := metricMap[opName]
					log.Infof("%s: total=%d,succ=%d,failed=%d\n", opName, metric.Total, metric.Succ, metric.Fail)
				}
				log.Info("\n")
			}
		}
	}(metricMap)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < *count; i++ {
				index := int(rand.Int31() % int32(len(opTypes)))
				opTypeIndex := opTypes[index]
				//	log.Info("begin index ", opTypeIndex, " run-", opMap[opTypeIndex])
				var ret int32
				switch opMap[opTypeIndex] {
				case "create_schema":
					ret = createSchmea(c, ctx, n)
				case "drop_schema":
					ret = dropSchmea(c, ctx, n)
				case "query_schema":
					ret = querySchmea(c, ctx, n)
				case "put_kv":
					ret = putKv(c, ctx, n)
				case "del_kv":
					ret = delKv(c, ctx, n)
				case "get_kv":
					ret = getKv(c, ctx, n)

				}
				atomic.AddUint64(&metricMap[opMap[opTypeIndex]].Total, 1)
				if ret < 0 {
					atomic.AddUint64(&metricMap[opMap[opTypeIndex]].Succ, 1)
				} else {
					atomic.AddUint64(&metricMap[opMap[opTypeIndex]].Fail, 1)
				}
			}
		case <-sigs:
			log.Info("stop benckmark....")
			return
		}
	}

}
