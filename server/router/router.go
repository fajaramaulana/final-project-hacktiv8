package router

import (
	"final-project/server/controllers"
	"final-project/server/middleware"

	"github.com/gin-gonic/gin"
)

type router struct {
	router *gin.Engine
	user   *controllers.UserController
	photo  *controllers.PhotoController
}

func NewRouter(user *controllers.UserController, photo *controllers.PhotoController) *router {
	return &router{
		router: gin.Default(),
		user:   user,
		photo:  photo,
	}
}

func (r *router) SetupRouter(port string) {
	user := r.router.Group("/users")
	user.POST("/register", r.user.Register)
	user.POST("/login", r.user.Login)
	user.PUT("/:userid", middleware.Authentication, r.user.Update)
	user.DELETE("", middleware.Authentication, r.user.Delete)

	photo := r.router.Group("/photos")
	photo.POST("", middleware.Authentication, r.photo.Create)

	r.router.Run(port)
}
