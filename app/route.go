package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sblog/core/dispatcher"
	server "sblog/core/gin"
)

var r = server.Route()

var diFront = dispatcher.Front()

//r.LoadHTMLGlob("templates/front/*")

var Index gin.HandlerFunc = func(c *gin.Context) {
	print("succuess")
	//c.JSON(200, gin.H{"iwas": "fine thank you"})
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Mainwebsite",
		"content": "Hello world",
	})
}

//===============================================================
//set app func
//===============================================================
var appSet gin.HandlerFunc = func(c *gin.Context) {
	r.LoadHTMLGlob("templates/front/*")
	print("ad")
}

func init() {
	diFront.Bind("tpl_front", appSet, 1)
	r.GET("/", diFront.Di(Index))
}
