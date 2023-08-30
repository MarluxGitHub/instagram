package interfaces

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/heartbeat/application"
	"github.com/MarluxGitHub/instagram/pkg/lib/api/rest"
	"github.com/gin-gonic/gin"
)

type HeartBeatRestEndpoint struct {
}

func NewHeartbeatRestEndpoint() rest.RestEndpoint {
	return &HeartBeatRestEndpoint{}
}

func (endpoint *HeartBeatRestEndpoint) SetRouting(engine *gin.Engine) {
	engine.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, application.GetHeartBeat())
	})
}