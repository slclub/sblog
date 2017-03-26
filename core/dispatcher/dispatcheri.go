package dispatcher

import (
	"github.com/gin-gonic/gin"
)

type Dispatcher interface {
	Di(c *gin.Context)
	Bind(c *gin.Context, pos int8)
}

//The location of inject function
const (
	//Execute before routing function.
	BEGIN int8 = 1 << iota
	//routing function.
	ROUTE_FUNC
	//Execute after routing function.
	END
)
