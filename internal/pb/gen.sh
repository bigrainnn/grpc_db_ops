





protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=dp_ops.yaml:. db_ops.proto
