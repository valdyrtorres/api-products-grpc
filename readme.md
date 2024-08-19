protoc src/proto/*.proto --go_out=./ --go-grpc_out=./

go get -v google.golang.org/protobuf
go get -v google.golang.org/grpc
