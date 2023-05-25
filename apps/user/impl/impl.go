package impl

import (
	"context"

	"github.com/cloudweops/phoenix/app"
	"github.com/cloudweops/phoenix/logger"
	"github.com/cloudweops/phoenix/logger/zap"
	"github.com/cloudweops/spark/apps/user"
	"github.com/cloudweops/spark/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var srv = &service{}

type service struct {
	log logger.Logger
	col *mongo.Collection
	user.UnimplementedRPCServer
}

func (s *service) Name() string {
	return user.AppName
}

func (s *service) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	uc := db.Collection(user.AppName)

	index := mongo.IndexModel{
		Keys:    bson.D{bson.E{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}
	_, err = uc.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		return err
	}

	s.col = uc
	s.log = zap.L().Named(user.AppName)
	return nil
}

func (s *service) Registry(server *grpc.Server) {
	user.RegisterRPCServer(server, srv)
}

func init() {
	app.RegistryInternalApp(srv)
	app.RegistryGRPCApp(srv)
}
