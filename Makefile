
.PHONY: protos init vendor

init : 
	go mod init github.com/AIF-user-system-management/entities

vendor : 
	go mod vendor

protos:
	protoc -I protos/ --go_out protos/  --go_opt paths=source_relative --go-grpc_out  protos/ --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative protos/user.proto 
