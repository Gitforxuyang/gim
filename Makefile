proto:
	protoc --proto_path=proto --go_out=proto/ proto/gim.proto

client:
	mkdir proto/im | true
	protoc --go_out=plugins=grpc:./proto/im --proto_path=proto  im.proto

.PHONY: proto