package router

import (
	"final-project/server/controllers"

	"github.com/gin-gonic/gin"
)

type router struct {
	router *gin.Engine
	user   *controllers.UserController
}

func NewRouter(user *controllers.UserController) *router {
	return &router{
		router: gin.Default(),
		user:   user,
	}
}

func (r *router) SetupRouter(port string) {
	v1 := r.router.Group("/api/v1")

	user := v1.Group("/users")
	user.POST("/register", r.user.Register)
	r.router.Run(port)
}
