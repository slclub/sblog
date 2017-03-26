package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sblog/core/dispatcher"
	server "sblog/core/gin"
)

var r = server.Route()

var diAdmin = dispatcher.Admin()

//r.LoadHTMLGlob("templates/front/*")

var Index gin.HandlerFunc = func(c *gin.Context) {
	print("succuess")
	//c.JSON(200, gin.H{"iwas": "fine thank you"})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Mainwebsite",
		"content": "Hello world",
	})
}

var adminSet gin.HandlerFunc = func(c *gin.Context) {
	r.LoadHTMLGlob("templates/admin/*")
	print("ad")
}

func init() {
	diAdmin.Bind("tmp_admin", adminSet, 1)
	r.GET("/sadmin/", diAdmin.Di(Index))
}
