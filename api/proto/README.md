Спецификация API gRPC: файл .proto
===

run: protoc --proto_path=./proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./transport \
	--go-grpc_out=./transport \
	./proto/*.proto