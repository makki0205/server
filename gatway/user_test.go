package gatway

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/makki0205/server/model"
	"github.com/makki0205/server/service"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	mock := service.SetMock(t)
	tests := httpTests{
		{
			name: "ok",
			args: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/v1/users", nil)
			},
			mock: func(mock service.Mock) {
				mock.User.EXPECT().GetAll().Return([]model.User{model.User{ID: "hoge"}}, nil)
			},
			want: http.StatusOK,
		},
		{
			name: "ng",
			args: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/v1/users", nil)
			},
			mock: func(mock service.Mock) {
				mock.User.EXPECT().GetAll().Return(nil, errMock)
			},
			want: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(mock)
			w := httptest.NewRecorder()
			NewGatway().ServeHTTP(w, tt.args())
			assert.Equal(t, w.Code, tt.want)
		})
	}
}
func TestCreateUsers(t *testing.T) {
	mock := service.SetMock(t)
	tests := httpTests{
		{
			name: "ok",
			args: func() *http.Request {
				req := model.User{}
				b, err := json.Marshal(req)
				assert.NoError(t, err)
				return httptest.NewRequest(http.MethodPost, "/v1/user", bytes.NewReader(b))
			},
			mock: func(mock service.Mock) {
				mock.User.EXPECT().Create(model.User{}).Return("id", nil)
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		tt.mock(mock)
		w := httptest.NewRecorder()
		NewGatway().ServeHTTP(w, tt.args())
		assert.Equal(t, w.Code, tt.want)
	}
}
