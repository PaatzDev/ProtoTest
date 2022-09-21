# GO PACKAGE

make in diffrent folder with go mod file
import it in proxy module by require and replace it

# PROTOC COMMANDS

run in working dir "prototest"

### FOR GO packages
protoc -I .\protos\ --go_out=gen/go --go_opt paths=source_relative  --go-grpc_out=gen/go --go-grpc_opt paths=source_relative .\protos\my_super_service.proto

## FOR Gateway
protoc -I .\protos\ --grpc-gateway_out=gen/go  --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true  .\protos\my_super_service.proto

### For OPENAPI 
 protoc -I . --openapiv2_out ./gen/openapiv2 --openapiv2_opt generate_unbound_methods=true 
 .\protos\my_super_service.proto

# Final step
register service in main method of proxy. Do that by calling the Register"service-name"HandlerFromEndpoint