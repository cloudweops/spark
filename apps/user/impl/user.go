package impl

import (
	"context"

	"github.com/cloudweops/phoenix/exception"
	"github.com/cloudweops/phoenix/pb/request"
	"github.com/cloudweops/spark/apps/user"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (s *service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	if err := s.save(ctx, u); err != nil {
		return nil, err
	}

	u.Password = ""
	return u, nil
}

func (s *service) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	filter := bson.M{}
	filter["username"] = req.Username
	ins := user.NewDefaultUser()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req.Username)
		}

		return nil, exception.NewInternalServerError("user %s error, %s", req.Username, err)
	}

	ins.Password = ""
	return ins, nil
}

func (s *service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	ins, err := s.DescribeUser(ctx, user.NewDescribeUserRequest(req.Username))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		err := ins.Update(req)
		if err != nil {
			return nil, err
		}
	case request.UpdateMode_PATCH:
		ins.Patch(req)
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
