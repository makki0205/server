package gatway

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/server/gatway/res"
	"github.com/makki0205/server/service"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("/users", GetUsers)
}

func GetUsers(c *gin.Context) {
	users, err := service.User.GetAll()
	if err != nil {
		fmt.Println("hoge")
		res.Internal(c)
		return
	}
	res.Json(users, c)
}
