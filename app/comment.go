package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var CommentAdd gin.HandlerFunc = func(c *gin.Context) {

	c.HTML(http.StatusOK, "forms.html", gin.H{
		"setting": Setting(""),
	})
}
