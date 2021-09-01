.PHONY: init
init:
	protoc -I proto/echo/ --go_out=client/generated --go-grpc_out=client/generated  service.proto && \
	protoc -I proto/echo/ --go_out=server/generated --go-grpc_out=server/generated  service.proto
