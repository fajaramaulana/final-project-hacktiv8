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

type SocmedController struct {
	socmedService *services.SocmedService
	userService   *services.UserService
}

func NewSocmedController(socmedService *services.SocmedService, userService *services.UserService) *SocmedController {
	return &SocmedController{socmedService: socmedService, userService: userService}
}

func (c *SocmedController) Create(ctx *gin.Context) {
	var req request.CreateSocialMedia

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

	data, err := c.socmedService.Create(&req, idUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, data)

}

func (c *SocmedController) Get(ctx *gin.Context) {
	email := ctx.GetString("email")

	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := c.socmedService.Get(idUser)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
		} else if err.Error() == "Social Media Not Found" {
			ctx.JSON(http.StatusNotFound, view.Error(http.StatusNotFound, err.Error()))
		} else {
			ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *SocmedController) Update(ctx *gin.Context) {
	var req request.UpdateSocialMedia

	idSocmed := ctx.Param("socialMediaId")

	socmedId, err := strconv.Atoi(idSocmed)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

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

	data, err := c.socmedService.Update(&req, idUser, socmedId)

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

func (c *SocmedController) Delete(ctx *gin.Context) {
	idSocmed := ctx.Param("socialMediaId")

	socmedId, err := strconv.Atoi(idSocmed)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	email := ctx.GetString("email")
	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := c.socmedService.Delete(idUser, socmedId)

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
