package main

import (
	"fmt"
	_ "sblog/admin"
	_ "sblog/app"
	server "sblog/core/gin"
)

var print = fmt.Println

func main() {
	r := server.Route()
	r.Run(":8001")
}
