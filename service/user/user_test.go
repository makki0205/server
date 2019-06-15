package user

import (
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
