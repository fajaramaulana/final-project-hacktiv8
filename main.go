package main

import (
	"final-project/config/db"
	"final-project/server/helper"
	"final-project/server/repositories/gorm"
	"fmt"
)

func main() {
	db, err := db.ConnectMysqlGorm()

	if err != nil {
		panic(err)
	}

	_ = gorm.NewUserRepository(db)
	_ = gorm.NewPhotoRepository(db)
	_ = gorm.NewCommentRepository(db)
	_ = gorm.NewSocialMediaRepository(db)

	token, err := helper.GenerateToken("fajar@gmail.com")
	fmt.Printf("%# v", token)
	// router := gin.Default()
}
