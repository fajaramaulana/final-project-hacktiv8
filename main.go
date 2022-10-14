package main

import (
	"final-project/config/db"
	"final-project/server/controllers"
	"final-project/server/repositories/gorm"
	"final-project/server/router"
	"final-project/server/services"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	db, err := db.ConnectMysqlGorm()

	if err != nil {
		panic(err)
	}

	userRepo := gorm.NewUserRepository(db)
	_ = gorm.NewPhotoRepository(db)
	_ = gorm.NewCommentRepository(db)
	_ = gorm.NewSocialMediaRepository(db)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	app := router.NewRouter(userController)

	err = godotenv.Load()

	if err != nil {
		panic(err)
	}

	app.SetupRouter(os.Getenv("PORT"))
}
