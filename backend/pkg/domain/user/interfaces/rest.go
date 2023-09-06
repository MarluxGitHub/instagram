package interfaces

import (
	"strconv"

	"github.com/MarluxGitHub/instagram/pkg/domain/user/application"
	domainmodels "github.com/MarluxGitHub/instagram/pkg/domain/user/domain/models"
	"github.com/MarluxGitHub/instagram/pkg/domain/user/interfaces/models"
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

		users := endpoint.application.GetUsers()

		response := []*models.UserResponse{}

		for _, user := range users {
			response = append(response, models.NewUserResponse(user))
		}

		c.JSON(200, response)
	})
	engine.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "Invalid id")
			return
		}
		c.JSON(200, models.NewUserResponse(endpoint.application.GetUser(id)))
	})

	engine.POST("/users", func(c *gin.Context) {
		user := &domainmodels.User{}
		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(400, "Invalid user")
			return
		}
		c.JSON(200, models.NewUserResponse(endpoint.application.CreateUser(user)))
	})

	engine.PUT("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "Invalid id")
			return
		}
		user := &domainmodels.User{}
		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(400, "Invalid user")
			return
		}
		user.ID = uint(id)
		c.JSON(200, models.NewUserResponse(endpoint.application.UpdateUser(user)))
	})

	engine.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "Invalid id")
			return
		}
		c.JSON(200, endpoint.application.DeleteUser(id))
	})

	engine.GET("/users/:id/followers", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "Invalid id")
			return
		}
		c.JSON(200, endpoint.application.GetFollowers(id))
	})
}