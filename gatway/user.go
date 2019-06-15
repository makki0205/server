package gatway

import (
	"github.com/gin-gonic/gin"
	"github.com/makki0205/server/gatway/res"
	"github.com/makki0205/server/model"
	"github.com/makki0205/server/service"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("/users", GetUsers)
	r.POST("/user", CreateUsers)
}

func GetUsers(c *gin.Context) {
	users, err := service.User.GetAll()
	if err != nil {
		res.Internal(c)
		return
	}
	res.Json(users, c)
}

func CreateUsers(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		res.BadRequest(err, c)
		return
	}
}
