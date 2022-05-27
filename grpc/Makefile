
proto-all: proto-grpc proto-http proto-gateway proto-openapi proto-doc

proto-grpc:
	protoc \
		--proto_path ./proto \
		--go_out=plugins=grpc:. \
		./proto/*.proto

proto-gateway:
	protoc \
		--proto_path ./proto \
		--grpc-gateway_out . \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt generate_unbound_methods=true \
        ./proto/*.proto

proto-http:
	protoc \
 		--proto_path ./proto \
 		--gohttp_out . \
 		./proto/*.proto

proto-openapi:
	protoc \
		--proto_path ./proto \
 		--openapiv2_out ./proto \
        --openapiv2_opt logtostderr=true \
        ./proto/*.proto

proto-doc:
	protoc \
		--proto_path ./proto \
		--doc_out ./proto \
		--doc_opt html,index.html \
		./proto/hello.proto
	protoc \
		--proto_path ./proto \
		--doc_out ./proto \
		--doc_opt markdown,hello.md \
		./proto/hello.proto
