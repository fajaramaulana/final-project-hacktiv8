package controllers

import (
	"final-project/server/controllers/view"
	"final-project/server/helper"
	"final-project/server/request"
	"final-project/server/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService *services.PhotoService
	userService  *services.UserService
}

func NewPhotoController(photoService *services.PhotoService, userService *services.UserService) *PhotoController {
	return &PhotoController{photoService: photoService, userService: userService}
}

func (c *PhotoController) Create(ctx *gin.Context) {
	var req request.CreatePhotoRequest
	email := ctx.GetString("email")

	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Validation", res))
		return
	}

	data, err := c.photoService.Create(&req, idUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

func (c *PhotoController) GetAll(ctx *gin.Context) {
	data, err := c.photoService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c PhotoController) Update(ctx *gin.Context) {
	var req request.UpdatePhotoRequest
	id := ctx.Param("photoid")
	email := ctx.GetString("email")

	idPhoto, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	userId, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Validation", res))
		return
	}

	data, err := c.photoService.Update(&req, idPhoto, userId)

	fmt.Printf("%# v", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, data)

}
