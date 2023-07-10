package routes

import (
	"github.com/MarluxGitHub/instagram/cmd/admin/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	r.GET("/heartbeat", controllers.Heartbeat)

}