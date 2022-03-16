go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/      
export GOPROXY=https://mirrors.aliyun.com/goproxy/  
go build -o  server server/server.go 
go build -o  client  client/client.go