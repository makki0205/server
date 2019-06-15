package gatway

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/makki0205/server/service"

	"github.com/gin-gonic/gin"
)

func TestNewGatway(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGatway(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGatway() = %v, want %v", got, tt.want)
			}
		})
	}
}

type httpTests []struct {
	name string
	args func() *http.Request
	mock func(mock service.Mock)
	want int
}

var errMock = errors.New("mock err")
