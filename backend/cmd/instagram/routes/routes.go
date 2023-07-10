package routes

import (
	"github.com/MarluxGitHub/instagram/cmd/instagram/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	r.GET("/heartbeat", controllers.Heartbeat)

}