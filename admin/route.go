package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sblog/core/dispatcher"
	server "sblog/core/gin"
)

var r = server.Route()
var diAdmin = dispatcher.Admin()

func init() {
	diAdmin.Bind("tmp_admin", adminSet, 1)
	r.Static("/static", ("./admin/static"))
	r.StaticFile("/jquery.js", "./static/jquery-3.2.0.min.js")
	r.Use(gin.Recovery())

	r.GET("/sadmin/", diAdmin.Di(Index))
	r.Any("/sadmin/post/save", diAdmin.Di(PostAdd))
}

var adminSet gin.HandlerFunc = func(c *gin.Context) {
	r.LoadHTMLGlob("templates/admin/*")
	//r.StaticFS("/static", http.Dir("./static/*filepath"))
}

var Index gin.HandlerFunc = func(c *gin.Context) {
	print("succuess")
	//c.JSON(200, gin.H{"iwas": "fine thank you"})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Sblog.admin",
		"content": "Hello world",
	})
}
