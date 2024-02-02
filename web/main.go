package main

import (
	"gin-sites/dao"
	"gin-sites/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// testRun()
	defer dao.Close()

	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// r.Static("/static", "./static")

	routes.RoutesInit(r)

	r.Run(":8080")
}
