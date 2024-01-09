goctl rpc -o demo.proto


goctl rpc protoc demo.proto --style=go_zero --go_out=. --go-grpc_out=. --zrpc_out=.