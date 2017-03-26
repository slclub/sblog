package gin

import (
	"github.com/gin-gonic/gin"
)

var ginn *gin.Engine

func Route() *gin.Engine {
	if ginn == nil {
		ginn = gin.Default()
	}
	return ginn
}
