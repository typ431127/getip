package main

import (
	"embed"
	"encoding/json"
	"flag"
	"getip/realip"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"strconv"
)

var port string

//go:embed templates
var tmpl embed.FS

func response(c *gin.Context) {
	c.Request.ParseForm()
	c.Request.ParseMultipartForm(33554432)
	response_code := 200
	format := c.DefaultQuery("format", "json")
	http_code := c.DefaultQuery("http_code", "200")
	if value, err := strconv.Atoi(http_code); err == nil {
		response_code = value
	}
	ip := c.ClientIP()
	djson := make(map[string]interface{})
	content_type := c.GetHeader("Content-Type")
	if err := c.ShouldBindJSON(&djson); err != nil && content_type == "application/json" {
		djson["message"] = "Submit json format error"
	}
	RealIp := realip.FromRequest(c.Request)
	response_json := make(map[string]interface{})
	response_json["ClientIp"] = ip
	response_json["RequestURI"] = c.Request.RequestURI
	response_json["Header"] = c.Request.Header
	response_json["Method"] = c.Request.Method
	response_json["RealIp"] = RealIp
	response_json["RequestJson"] = djson
	response_json["RequestPostForm"] = c.Request.PostForm
	response_json["Response_code"] = response_code
	response_json["Content-Type"] = c.Request.Header.Get("Content-Type")
	bytejson, _ := json.MarshalIndent(&djson, "", "  ")
	log.Printf("\n============================================================================\n"+
		"Header:%s\n"+
		"IP:%s\n"+
		"X-Forwarded-For:%s\n"+
		"X-Real-Ip:%s\n"+
		"X-Forwarded-Host:%s\n"+
		"RemoteAddr:%s\n"+
		"Content-Type:%s\n"+
		"RequestJson::%s\n"+
		"RequestPostForm::%s\n",
		c.Request.Header,
		c.ClientIP(),
		c.Request.Header.Get("X-Forwarded-For"),
		c.Request.Header.Get("X-Real-Ip"),
		c.Request.Header.Get("X-Forwarded-Host:"),
		c.Request.RemoteAddr,
		c.Request.Header.Get("Content-Type"),
		string(bytejson),
		c.Request.PostForm)
	if format == "json" {
		c.JSON(response_code, response_json)
	} else {
		c.HTML(response_code, "index.tmpl", gin.H{
			"response_json": response_json,
			"Header":        c.Request.Header,
		})
	}
}

func init() {
	flag.StringVar(&port, "port", ":8080", "端口")
}
func main() {
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	t, _ := template.ParseFS(tmpl, "templates/*.tmpl")
	r.SetHTMLTemplate(t)
	v1 := r.Group("/")
	v1.GET("/*router", response)
	v1.HEAD("/*router", response)
	v1.POST("/*router", response)
	v1.PUT("/*router", response)
	v1.DELETE("/*router", response)
	v1.OPTIONS("/*router", response)
	r.Run(port)
}
