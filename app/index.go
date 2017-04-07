package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sblog/source"
	"strconv"
)

var Index gin.HandlerFunc = func(c *gin.Context) {
	//c.JSON(200, gin.H{"iwas": "fine thank you"})
	post := source.NewPost()

	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	Print(uint(page))
	post.Limit(0, 20)
	post.Page(uint(page))
	ret := post.Find(post, "", []interface{}{})

	//Print(ret)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"web_title": "Aixgl.艾辛阁",
		"content":   "Hello world",
		"postList":  ret,
		"prepage":   page - 1,
		"nextpage":  page + 1,
	})
}

func init() {
	//r.Any("/", diFront.Di(index))
}

var Page = func() {
}
