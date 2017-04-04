package admin

import (
	"github.com/gin-gonic/gin"
	"sblog/source"
	"strconv"
)

var PostAdd gin.HandlerFunc = func(c *gin.Context) {
	m := make(map[string]interface{})
	m["p_id"] = c.PostForm("ID")
	m["c_id"] = c.PostForm("c_id")
	m["uid"] = c.PostForm("uid")
	m["content"] = c.PostForm("content")
	m["title"] = c.PostForm("title")
	m["tags"] = c.PostForm("tags")
	m["sort"] = c.PostForm("sort")
	post := source.NewPost()
	post.Assign(m)
	post.Save(post)

	//ret := post.Exists(post)
	c.JSON(200, m)
}

var PostFind gin.HandlerFunc = func(c *gin.Context) {
	post := source.NewPost()

	v, _ := strconv.Atoi(c.Query("page"))
	post.Page(uint(v))
	ret := post.Find(post, "", []interface{}{})
	c.JSON(200, (ret))
}

var PostAddHtml gin.HandlerFunc = func(c *gin.Context) {
	post := source.NewPost()

	v, _ := strconv.Atoi(c.Query("ID"))
	post.ID(v)
	ret := post.Exists(post)
	c.JSON(200, (ret))
}
