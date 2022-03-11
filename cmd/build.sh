go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/      
export GOPROXY=https://mirrors.aliyun.com/goproxy/  
go build -o  test_server server/server.go 
go build -o  test_client  client/client.go