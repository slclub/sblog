package app

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	sbtext "sblog/provider/text"
	"sblog/source"
	"strconv"
)

var PostFind gin.HandlerFunc = func(c *gin.Context) {
}

var PostDetail gin.HandlerFunc = func(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("ID"))

	post := source.NewPost()
	ret := post.FindOne(post, id)

	ct := ret["content"].(string)
	ct = sbtext.DeScript(ct + "<script>alert('nothing')</script>")
	ret["content"] = template.HTML(ct)
	c.HTML(http.StatusOK, "detail.tmpl", gin.H{
		"setting": Setting(""),
		"detail":  ret,
	})

}
