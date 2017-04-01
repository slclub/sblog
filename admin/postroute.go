package admin

import (
	"github.com/gin-gonic/gin"
	"sblog/source"
	"strconv"
)

var PostAdd gin.HandlerFunc = func(c *gin.Context) {
	m := make(map[string]interface{})
	m["p_id"] = c.Query("p_id")
	m["c_id"] = c.Query("c_id")
	m["uid"] = c.Query("uid")
	m["content"] = c.Query("content")
	m["title"] = c.Query("title")
	post := source.NewPost()
	post.Assign(m)
	post.Save(post)
}

var PostFind gin.HandlerFunc = func(c *gin.Context) {
	post := source.NewPost()

	v, _ := strconv.Atoi(c.Query("page"))
	post.Page(uint(v))
	ret := post.Find(post, "", []interface{}{})
	c.JSON(200, (ret))
}
