package controller

import (
	"github.com/gin-gonic/gin"
	"jmg-my-instagram-ms/services"
	"net/http"
)

type UserInstagramController struct {
	userInstaService services.IUserService
}

func (c UserInstagramController) Search(ctx *gin.Context) {
	users, err := c.userInstaService.SearchUsers(ctx.Query("name"), 10)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c UserInstagramController) GetInfoProfile(context *gin.Context) {

}

func NewUserInstagramController(userInstaService services.IUserService) UserInstagramController {
	return UserInstagramController{userInstaService}
}
