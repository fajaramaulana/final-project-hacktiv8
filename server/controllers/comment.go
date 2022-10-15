package controllers

import (
	"final-project/server/controllers/view"
	"final-project/server/helper"
	"final-project/server/request"
	"final-project/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentService
	userService    *services.UserService
	photoService   *services.PhotoService
}

func NewCommentController(commentService *services.CommentService, userService *services.UserService, photoService *services.PhotoService) *CommentController {
	return &CommentController{commentService: commentService, userService: userService, photoService: photoService}
}

func (c *CommentController) Create(ctx *gin.Context) {
	var req request.CreateCommentRequest

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

	checkIfPhotoIdExist, err := c.photoService.GetPhotoById(req.PhotoId)

	if !checkIfPhotoIdExist {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, "Photo Id Not Found"))
		return
	}

	data, err := c.commentService.Create(idUser, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, data)
}
