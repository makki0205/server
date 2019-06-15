package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(code int, c *gin.Context) {
	c.Status(code)
}

func Internal(c *gin.Context) {
	Status(http.StatusInternalServerError, c)
}

func Json(json interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, json)
}
