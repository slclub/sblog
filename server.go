package main

import (
	"fmt"
	_ "sblog/admin"
	_ "sblog/app"
	"sblog/core/gin"
)

var print = fmt.Println

func main() {
	r := gin.Route()
	r.Run(":8001")
}
