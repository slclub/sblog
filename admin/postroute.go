package admin

import (
	"github.com/gin-gonic/gin"
	"sblog/core/model"
	"sblog/source"
)

var PostAdd gin.HandlerFunc = func(c *gin.Context) {
	m := make(map[string]interface{})
	m["p_id"] = c.Query("p_id")
	m["c_id"] = c.Query("c_id")
	m["uid"] = c.Query("uid")
	post := &source.Post{}
	post.Model = &model.Model{}
	post.Assign(m)
	post.Save(post)
}
