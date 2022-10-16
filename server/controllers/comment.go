package controllers

import (
	"final-project/server/controllers/view"
	"final-project/server/helper"
	"final-project/server/request"
	"final-project/server/services"
	"net/http"
	"strconv"

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
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

func (c *CommentController) GetAll(ctx *gin.Context) {
	email := ctx.GetString("email")
	idUser, err := c.userService.GetUserIdByEmail(email)

	data, err := c.commentService.GetAll(idUser)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
		} else if err.Error() == "Comment Not Found" {
			ctx.JSON(http.StatusNotFound, view.Error(http.StatusNotFound, err.Error()))
		} else {
			ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *CommentController) Update(ctx *gin.Context) {
	var req request.UpdateCommentRequest
	idComment := ctx.Param("commentid")

	commentId, err := strconv.Atoi(idComment)

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

	data, err := c.commentService.Update(idUser, commentId, &req)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, data)

}

func (c *CommentController) Delete(ctx *gin.Context) {
	idComment := ctx.Param("commentid")

	commentId, err := strconv.Atoi(idComment)

	email := ctx.GetString("email")

	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	err = c.commentService.Delete(idUser, commentId)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, view.ResponseDeleteComment{
		Message: "Your comment has been successfully deleted",
	})
}
