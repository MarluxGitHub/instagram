package interfaces

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/rss/application"
	"github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/models"
	"github.com/MarluxGitHub/instagram/pkg/lib/api/rest"
	"github.com/gin-gonic/gin"
)

type RssEndpoint struct {
	application application.RssApplication
}

func NewRssEndpoint(c models.Configuration) rest.RestEndpoint {
	return &RssEndpoint{
		application: application.NewRssApplication(c),
	}
}

func (endpoint *RssEndpoint) SetRouting(engine *gin.Engine) {
	engine.GET("/rss", func(c *gin.Context) {
		c.JSON(200, endpoint.application.GetAllFeeds())
	})
}