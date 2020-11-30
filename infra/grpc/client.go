package grpc

import (
	"gim/conf"
	"gim/infra/utils"
	"gim/proto/im"
	"google.golang.org/grpc"
)

func InitClient(config *conf.Config) im.ImClient {
	conn, err := grpc.Dial(config.GRpc.Endpoint,
		grpc.WithInsecure())
	utils.Must(err)
	client := im.NewImClient(conn)
	return client
}
