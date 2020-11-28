proto:
	protoc --proto_path=proto --go_out=proto/ proto/gim.proto

.PHONY: proto