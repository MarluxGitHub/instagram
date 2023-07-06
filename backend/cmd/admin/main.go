package main

import (
	"github.com/MarluxGitHub/instagram/cmd/admin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Setup(r)

	r.Run() // listen and serve on
}