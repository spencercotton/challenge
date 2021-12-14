.PHONY: gen-go

gen-go:
	protoc --go_out=paths=source_relative:./internal/grpc/pb --go-grpc_out=paths=source_relative:./internal/grpc/pb --go-grpc_opt=paths=source_relative service.proto

start:
	docker build -t coupons . && docker-compose up