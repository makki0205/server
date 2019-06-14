package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/makki0205/server/service/user/mock"
)

type Mock struct {
	User   *MockUserService
	Finish func()
}

func SetMock(t *testing.T) Mock {
	c := gomock.NewController(t)

	mock := Mock{
		User:   NewMockUserService(c),
		Finish: c.Finish,
	}
	User = mock.User
	return mock
}
