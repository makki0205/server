package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/makki0205/server/model"
)

func TestSetMock(t *testing.T) {
	m := SetMock(t)
	users := []model.User{
		{
			ID:       "id",
			Email:    "email",
			Password: "pass",
		},
	}
	m.User.EXPECT().GetAll().Return(users, nil)
	res, err := User.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, res, users)
	m.User.EXPECT().GetAll().Return(nil, errors.New("エラー"))
	res, err = User.GetAll()
	assert.Error(t, err)
	assert.Nil(t, res)
}
