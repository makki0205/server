package gatway

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewGatway() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	userRouter(v1)
	r.Use(cros)
	r.GET("/test", func(c *gin.Context) {
		c.Writer.Header().Add("Set-Cookie", "CloudFront-Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiaHR0cHM6Ly9vbmRlbWFuZC1kZXYuY2EtbG92ZS5pby8qIiwiQ29uZGl0aW9uIjp7IkRhdGVMZXNzVGhhbiI6eyJBV1M6RXBvY2hUaW1lIjoxNTkyMjg3MDMwfX19XX0_; Path=/license/; Domain=ngrok.io; HttpOnly")
		c.Writer.Header().Add("Set-Cookie", "CloudFront-Signature=EyvOKMPFwjoVEdgD87HPuqV0-RJEGnUofsCIRUp7gNhzNKGvJ7Nk2XBN40F~3sGI3MsSoG8NBWOIG2DNFNiPYhCH8NDCRlnMJIKpAsfNKqviSyrzJ5Z-VE17rqhA45V3YDcFZxs-uHhfGW92uHvxdJXszVRfZY~3gUIkbIwUcpvtv~Y~1bUQ7-N3HN1dyk8MUunFkLx-K-jF8Gkv9Gc2U2WfL2gA4AUFOxUch8fpTu5vwwqsM3AFwdMSpSSuaFv58VOgOXVXo7VKNoyR8pqcUsJ3Z2NXUy1quz1KqoJRqpzHGfAuecih3h1JwdULX6W-FcwhP6UC569sq0TVYzl2tw__; Path=/license/; Domain=ca-love.io; HttpOnly")
		c.Writer.Header().Add("Set-Cookie", "CloudFront-Key-Pair-Id=APKAJQXAGHOH2DN6DDVQ; Path=/license/; Domain=ca-love.io")
		c.SetCookie("hoge", "hoge", 0, "/", "", false, false)
		c.JSON(http.StatusOK, "hoge")
	})
	r.GET("/license/hoge", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
	})
	return r
}

func cros(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	c.Header("Access-Control-Allow-Credentials", "true")

	// 任意のheaderをclientに向けてexposeする
	c.Header("Access-Control-Expose-Headers", "Set-Cookie")
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Set("start_time", time.Now())
	c.Next()

}
