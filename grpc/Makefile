# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

#编译hello_http.proto为hello_http.pb.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=grpc-hello-world/proto/google/api:. ./hello.proto

#编译hello_http.proto为hello_http.pb.gw.proto
protoc --grpc-gateway_out=logtostderr=true:.. ./hello.proto

# 编译生成 *pb.go (内带grpc编码)
proto-grpc:
	protoc \
		--proto_path=. \
		--go_out=plugins=grpc:. \
		./proto/*.proto

proto-http:
	protoc \
 		--proto_path=. \
 		--go_out=plugins=http:. \
 		./proto/*.proto

proto-all:
	protoc \
		--proto_path=. \
        --govalidators_out=. \
        --go_out=plugins=grpc:.\
        ./proto/*.proto