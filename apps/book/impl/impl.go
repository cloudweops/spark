package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/cloudweops/phoenix/app"
	"github.com/cloudweops/phoenix/logger"
	"github.com/cloudweops/phoenix/logger/zap"
	"google.golang.org/grpc"

	"github.com/cloudweops/spark/apps/book"
	"github.com/cloudweops/spark/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	book.UnimplementedServiceServer
}

func (s *service) Config() error {

	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return book.AppName
}

func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGRPCApp(svr)
}
