package rest

import "github.com/gin-gonic/gin"

type RestEndpoint interface {
	SetRouting(*gin.Engine)
}