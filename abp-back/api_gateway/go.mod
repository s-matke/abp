module github.com/s-matke/abp/abp-back/api_gateway

go 1.20

replace github.com/s-matke/abp/abp-back/common => ../common

require (
	github.com/gorilla/handlers v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/s-matke/abp/abp-back/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.55.0
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
