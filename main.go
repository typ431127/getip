package main

import (
	"flag"
	"getip/realip"
	"github.com/gin-gonic/gin"
)

var port string

func response(c *gin.Context) {
	ip := c.ClientIP()
	RealIp := realip.FromRequest(c.Request)
	response_json := make(map[string]interface{})
	response_json["Client-Ip"] = ip
	response_json["RequestURI"] = c.Request.RequestURI
	response_json["Header"] = c.Request.Header
	response_json["Method"] = c.Request.Method
	response_json["RealIp"] = RealIp
	c.JSON(200, response_json)
}

func init() {
	flag.StringVar(&port, "port", ":8080", "端口")
}
func main() {
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("/*action", response)
		v1.HEAD("/*action", response)
		v1.POST("/*action", response)
		v1.PUT("/*action", response)
		v1.DELETE("/*action", response)
		v1.OPTIONS("/*action", response)
	}
	r.Run(port)
}
