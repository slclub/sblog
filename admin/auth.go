package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sblog/core"
	"sblog/provider/jwt"
	"strings"
	"time"
)

var AuthFunc gin.HandlerFunc = func(c *gin.Context) {
	tokenAr := c.Request.Header["Token"]
	token := strings.Join(tokenAr, ".")
	tokenMap, _ := jwt.Decode(token)

	if tokenMap["valid"] != nil {
		c.Redirect(http.StatusMovedPermanently, "/sadmin/notvalidtoken")
	}
}

var NotValidToken gin.HandlerFunc = func(c *gin.Context) {
	tokenAr := c.Request.Header["Token"]
	token := strings.Join(tokenAr, ".")
	tokenMap, _ := jwt.Decode(token)

	exp := tokenMap["exp"].(int64)

	exp = exp - (time.Now().Unix())
	Print(tokenMap, len(token), exp)
	if tokenMap["valid"] != nil {
		valid := tokenMap["valid"].(error)
		err := errors.New("token is expired or not valid token :" + valid.Error())
		result := core.NewResult(token, err)
		Print(err.Error())
		c.JSON(200, result)
	}
}
