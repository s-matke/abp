module github.com/s-matke/abp/abp-back/user_service

go 1.20

replace github.com/s-matke/abp/abp-back/common => ../common

require (
	github.com/google/uuid v1.3.0
	github.com/s-matke/abp/abp-back/common v0.0.0-20230422221353-aac738e9c5a1
	google.golang.org/grpc v1.55.0
	gorm.io/driver/postgres v1.5.0
	gorm.io/gorm v1.25.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
