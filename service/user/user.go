package user

//go:generate mkdir -p mock
//go:generate mockgen -source=./user.go -package=mock -destination=./mock/mock.go

import (
	"github.com/makki0205/server/model"
	"github.com/rs/xid"
)

type UserService interface {
	Create(user model.User) (string, error)
	GetAll() ([]model.User, error)
}

type userService struct {
	user []model.User
}

func NewUserService() *userService {
	return &userService{
		user: make([]model.User, 0),
	}
}

func (u *userService) Create(user model.User) (string, error) {
	user.ID = xid.New().String()
	u.user = append(u.user, user)
	return user.ID, nil
}

func (u *userService) GetAll() ([]model.User, error) {
	return u.user, nil
}
