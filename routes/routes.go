package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tharaka911/go-redis-api/controllers"
	"os"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	//need to create / as a health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome lakshan-test-api,seems healthy",
		})
	})

	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.PostGetAll)
	r.GET("/posts/:id", controllers.PostGet)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	return r

}
