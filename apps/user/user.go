package user

import (
	"time"

	"github.com/cloudweops/phoenix/exception"
	"github.com/cloudweops/phoenix/pb/request"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

const AppName = "user"

var (
	validate = validator.New()
)

func New(req *CreateUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	hashPassword, err := NewHashedPassword(req.User.Password)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	u := &User{
		Username:    req.User.Username,
		Password:    hashPassword,
		Email:       req.User.Email,
		PhoneNumber: req.User.PhoneNumber,
		Address:     req.User.Address,
		CreatedAt:   time.Now().Unix(),
	}

	return u, nil
}

func NewDefaultUser() *User {
	return &User{}
}

func NewHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// NewCreateUserRequest 创建请求
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		User: &User{},
	}
}

func NewDescribeUserRequest(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		Username: username,
	}
}

func NewUpdateUserRequestWithPut() *UpdateUserRequest {
	return &UpdateUserRequest{
		UpdateMode: request.UpdateMode_PUT,
	}
}

func NewUpdateUserRequestWithPatch() *UpdateUserRequest {
	return &UpdateUserRequest{
		UpdateMode: request.UpdateMode_PATCH,
	}
}

// Validate 校验请求是否合法
func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

// Validate 校验请求是否合法
func (req *DescribeUserRequest) Validate() error {
	return validate.Struct(req)
}

// Validate 校验请求是否合法
func (req *UpdateUserRequest) Validate() error {
	return validate.Struct(req)
}

// Update
func (u *User) Update(req *UpdateUserRequest) error {
	// 登录时需要更新 login_ua login_ip login_at
	// 退出登录需要更新 logout_at
	// 更新用户属性需要更新 update_at
	hashPassword, err := NewHashedPassword(u.Password)
	u.Password = hashPassword
	if err != nil {
		return err
	}

	u.Email = req.Email
	u.PhoneNumber = req.PhoneNumber
	u.Address = req.Address
	u.UpdateAt = time.Now().Unix()
	return nil
}

func (u *User) Patch(req *UpdateUserRequest) error {
	// 更新用户属性需要更新 update_at
	u.UpdateAt = time.Now().Unix()
	if req.Password != "" {
		hashPassword, err := NewHashedPassword(u.Password)
		u.Password = hashPassword
		if err != nil {
			return err
		}
	}

	if req.Email != "" {
		u.Email = req.Email
	}
	if req.PhoneNumber != "" {
		u.PhoneNumber = req.PhoneNumber
	}
	if req.Address != "" {
		u.Address = req.Address
	}
	return nil
}
