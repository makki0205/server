package user

//go:generate mkdir -p mock
//go:generate mockgen -source=./user.go -package=mock -destination=./mock/mock.go

import "github.com/makki0205/server/model"

type User interface {
	Create(user model.User) (string, error)
	Update(user model.User) (model.User, error)
}

type user struct {
}
