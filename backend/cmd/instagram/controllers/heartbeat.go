package controllers

import "github.com/gin-gonic/gin"

func Heartbeat(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "alive",
	})
}