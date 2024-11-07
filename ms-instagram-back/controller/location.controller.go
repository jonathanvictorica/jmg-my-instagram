package controller

import (
	"github.com/gin-gonic/gin"
	"jmg-my-instagram-ms/services"
	"net/http"
)

type LocationController struct {
	locationService services.ILocationService
}

func NewLocationController(locationService services.ILocationService) LocationController {
	return LocationController{locationService: locationService}
}

func (c *LocationController) Search(ctx *gin.Context) {
	locations, err := c.locationService.SearchLocations(ctx.Query("name"), 10)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, locations)
}
