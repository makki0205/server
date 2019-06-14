package gatway

import "github.com/gin-gonic/gin"

func NewGatway() *gin.Engine {
	r := gin.New()
	return r
}
