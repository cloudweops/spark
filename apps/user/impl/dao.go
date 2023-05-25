package impl

import (
	"context"

	"github.com/cloudweops/phoenix/exception"
	"github.com/cloudweops/spark/apps/user"
	"gopkg.in/mgo.v2/bson"
)

func (s *service) save(ctx context.Context, u *user.User) error {
	if _, err := s.col.InsertOne(ctx, u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Username, err)
	}

	return nil
}

func (s *service) update(ctx context.Context, ins *user.User) error {
	filter := bson.M{}
	filter["username"] = ins.Username
	if _, err := s.col.UpdateOne(ctx, filter, bson.M{"$set": ins}); err != nil {
		return exception.NewInternalServerError("update user(%s) document error, %s",
			ins.Username, err)
	}

	return nil
}
