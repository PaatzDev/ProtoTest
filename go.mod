module grpc-rest-proxy

go 1.17

require (
	github.com/golang/glog v1.0.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

require skylinx.aero/airline v0.0.0

replace skylinx.aero/airline v0.0.0 => ./airline
