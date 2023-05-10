module github.com/mihajlo-ra92/XML/auth_service

go 1.20

replace github.com/mihajlo-ra92/XML/common => ../common

require google.golang.org/grpc v1.54.0

require github.com/mihajlo-ra92/XML/common v0.0.0-20230503231415-6b0d7acff76d

require (
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
