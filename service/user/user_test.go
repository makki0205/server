package user

import (
	"reflect"
	"testing"

	"github.com/makki0205/server/model"
)

func Test_userService_Create(t *testing.T) {
	type fields struct {
		user  []model.User
		newID func() string
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				newID: func() string {
					return "id"
				},
			},
			want: "id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				user:  tt.fields.user,
				newID: tt.fields.newID,
			}
			got, err := u.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewUserService()
		})
	}
}

func Test_userService_GetAll(t *testing.T) {
	type fields struct {
		user  []model.User
		newID func() string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.User
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				user: []model.User{
					{
						ID:       "hoge",
						Email:    "email",
						Password: "pass",
					},
				},
			},
			want: []model.User{
				{
					ID:       "hoge",
					Email:    "email",
					Password: "pass",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				user:  tt.fields.user,
				newID: tt.fields.newID,
			}
			got, err := u.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
