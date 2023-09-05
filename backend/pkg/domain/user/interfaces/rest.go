package interfaces

import (
	"strconv"

	"github.com/MarluxGitHub/instagram/pkg/domain/user/application"
	"github.com/MarluxGitHub/instagram/pkg/lib/api/rest"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserEndPoint struct {
	application application.UserApplication
}

func NewUserEndPoint(db *gorm.DB) rest.RestEndpoint {
	return &UserEndPoint{
		application: application.NewUserApplication(db),
	}
}

func (endpoint *UserEndPoint) SetRouting(engine *gin.Engine) {
	engine.GET("/users", func(c *gin.Context) {
		c.JSON(200, endpoint.application.GetUsers())
	})
	engine.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "Invalid id")
			return
		}
		c.JSON(200, endpoint.application.GetUser(id))
	})
}