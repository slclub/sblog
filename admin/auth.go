package admin

import (
	"github.com/gin-gonic/gin"
)

var AuthFunc gin.HandlerFunc = func(c *gin.Context) {
	Print("TOKEN:=")
	Print(c.Request.Header["Token"])
	Print(c.Request.URL.Path)
}
