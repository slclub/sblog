package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sblog/core/dispatcher"
)

var print = fmt.Println

func main() {
	r := gin.Default()
	di := dispatcher.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	r.Run(":8001")
}
