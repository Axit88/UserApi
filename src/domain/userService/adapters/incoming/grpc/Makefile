create:
	protoc --proto_path=proto proto/*.proto --go_out=pb/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=pb/

clean:
	rm pb/*.go
	