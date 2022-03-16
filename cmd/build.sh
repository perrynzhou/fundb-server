go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/      
export GOPROXY=https://mirrors.aliyun.com/goproxy/  
go build -o  dbserver server/server.go 
go build -o  dbclient  client/client.go