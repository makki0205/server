package gatway

import "github.com/gin-gonic/gin"

func NewGatway() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1")
	userRouter(v1)
	return r
}
