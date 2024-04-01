package main

import (
	"embed"
	"encoding/json"
	"flag"
	"getip/realip"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
	"strconv"
)

var port string
var Hostname string
var ShowHostname bool

//go:embed templates
var tmpl embed.FS

func response(c *gin.Context) {
	c.Request.ParseForm()
	c.Request.ParseMultipartForm(33554432)
	responseCode := 200
	format := c.DefaultQuery("format", "json")
	httpCode := c.DefaultQuery("http_code", "200")
	if value, err := strconv.Atoi(httpCode); err == nil {
		responseCode = value
	}
	ip := c.ClientIP()
	djson := make(map[string]interface{})
	contentType := c.GetHeader("Content-Type")
	if err := c.ShouldBindJSON(&djson); err != nil && contentType == "application/json" {
		djson["message"] = "Submit json format error"
	}
	RealIp := realip.FromRequest(c.Request)
	responseJson := make(map[string]interface{})
	responseJson["ClientIp"] = ip
	responseJson["RequestURI"] = c.Request.RequestURI
	responseJson["Header"] = c.Request.Header
	responseJson["Method"] = c.Request.Method
	responseJson["RealIp"] = RealIp
	responseJson["RequestJson"] = djson
	responseJson["RequestPostForm"] = c.Request.PostForm
	responseJson["Response_code"] = responseCode
	responseJson["Content-Type"] = c.Request.Header.Get("Content-Type")
	//
	if ShowHostname {
		responseJson["Hostname"] = Hostname
	}
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
		c.JSON(responseCode, responseJson)
	} else {
		c.HTML(responseCode, "index.tmpl", gin.H{
			"response_json": responseJson,
			"Header":        c.Request.Header,
		})
	}
}

func init() {
	flag.BoolVar(&ShowHostname, "hostname", false, "返回json中显示主机名")
	flag.StringVar(&port, "port", ":8080", "端口")
}
func main() {
	flag.Parse()
	if ShowHostname {
		Hostname, _ = os.Hostname()
	}
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
