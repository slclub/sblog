package admin

import (
	"github.com/gin-gonic/gin"
	"sblog/core"
	"sblog/provider/jwt"
	"sblog/provider/md5"
	"sblog/source"
)

func init() {
	diAdmin.NotAllow("/sadmin/s-lg", []string{"jwt_token"})
	r.GET("/sadmin/s-lg", diAdmin.Di(UserLogin))
}

var UserLogin gin.HandlerFunc = func(c *gin.Context) {
	username := c.Query("lg-name")
	password := c.Query("lg-pwd")
	m := make(map[string]interface{})

	password = md5.Md5(&password)

	m["username"] = username
	m["password"] = password

	admin := source.NewAdmin()
	ret := admin.FindOne(admin, " username=? AND password =? ", []interface{}{username, password})

	if ret != nil {
		delete(ret, "password")
		token, err := jwt.Encode(ret)
		if err == nil {
			ret["token"] = token
			ret["untoken"], _ = jwt.Decode(token)
		}
	}

	var result *core.Result
	result = core.NewResult(ret, nil)
	c.JSON(200, result)
}
