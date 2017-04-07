package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sblog/core/dispatcher"
	server "sblog/core/gin"
)

var r = server.Route()

var diFront = dispatcher.Front()
var Print = fmt.Println

//r.LoadHTMLGlob("templates/front/*")

//===============================================================
//set app func
//===============================================================
var appSet gin.HandlerFunc = func(c *gin.Context) {
	r.LoadHTMLGlob("templates/front/*.tmpl")
}

func init() {
	diFront.Bind("tpl_front", appSet, 1)
	r.GET("/", diFront.Di(Index))
	r.Static("/front", ("./static"))
	r.GET("/v", diFront.Di(PostDetail))
	//在后台路由设置中已经引入js路径了
	//r.StaticFile("/jquery.js", "./static/jquery-3.2.0.min.js")
	//r.StaticFile("/jquery.cookie.js", "./static/jquery-cookie/src/jquery.cookie.js")
	r.Use(gin.Recovery())

}
