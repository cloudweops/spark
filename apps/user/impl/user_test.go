package impl_test

import (
	"testing"

	"github.com/cloudweops/spark/apps/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserSuccess(t *testing.T) {
	should := assert.New(t)
	req := user.NewCreateUserRequest()
	req.User.Username = "admin"
	req.User.Password = "12346"
	req.User.Email = "moweilong@utopa.com.cn"
	req.User.PhoneNumber = "12345678901"
	_, err := impl.CreateUser(ctx, req)
	should.NoError(err)
}

func TestCreateUserFailedWithUsername(t *testing.T) {
	should := assert.New(t)
	req := user.NewCreateUserRequest()
	req.User.Username = "admi"
	req.User.Password = "12346"
	req.User.Email = "moweilong@utopa.com.cn"
	req.User.PhoneNumber = "12345678901"
	_, err := impl.CreateUser(ctx, req)
	t.Log(err)
	should.Contains(err.Error(), "Username")

}

func TestDescribeUserSuccess(t *testing.T) {
	should := assert.New(t)
	req := user.NewDescribeUserRequest("admin")
	u, err := impl.DescribeUser(ctx, req)
	should.NoError(err)
	t.Log(u)
}

func TestDescribeUserFailed(t *testing.T) {
	should := assert.New(t)
	req := user.NewDescribeUserRequest("admin1")
	u, err := impl.DescribeUser(ctx, req)
	should.Contains(err.Error(), "user admin1 not found")
	t.Log(u)
}

func TestUpdateUserWithPut(t *testing.T) {
	should := assert.New(t)
	req := user.NewUpdateUserRequestWithPut()
	req.Username = "admin"
	req.Password = "1234567"
	req.Email = "18550039021@163.com"
	req.PhoneNumber = "18550039021"
	req.Address = "gz"
	u, err := impl.UpdateUser(ctx, req)
	should.NoError(err)
	t.Log(u)

}

func TestUpdateUserWithPatch(t *testing.T) {
	should := assert.New(t)
	req := user.NewUpdateUserRequestWithPatch()
	req.Username = "admin"
	req.Email = "185@163.com"
	req.Address = "gdgz"
	u, err := impl.UpdateUser(ctx, req)
	should.NoError(err)
	t.Log(u)

}
