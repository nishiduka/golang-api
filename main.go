package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nishiduka/rest-api/controllers"
	"github.com/nishiduka/rest-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsList)
	r.GET("/posts/:id", controllers.PostDetails)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.POST("/posts", controllers.PostCreate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
