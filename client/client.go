package client

import (
	"github.com/cloudweops/phoenix/logger"
	"github.com/cloudweops/phoenix/logger/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cloudweops/spark/apps/book"
)

var (
	client *ClientSet
)

// SetGlobal todo
func SetGlobal(cli *ClientSet) {
	client = cli
}

// C Global
func C() *ClientSet {
	return client
}

// Client 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Book服务的SDK
func (c *ClientSet) Book() book.ServiceClient {
	return book.NewServiceClient(c.conn)
}

func NewClient() (*ClientSet, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(
		":8089",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conn: conn,
		log:  log,
	}, nil
}
