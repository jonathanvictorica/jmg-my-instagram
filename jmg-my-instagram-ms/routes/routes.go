package routes

import (
	"github.com/gin-gonic/gin"
	"jmg-my-instagram-ms/factory"
)

func SetupRouter(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	v1Locations := v1.Group("/locations")
	v1Posts := v1.Group("/posts")
	v1Users := v1.Group("/users")
	v1History := v1.Group("/history")
	v1Histories := v1.Group("/histories")
	v1HistoryContents := v1.Group("/history-contents")

	v1Locations.GET("/search", factory.FactoryContainer().LocationController.Search)

	v1Posts.GET("/get-by-user/:userId", factory.FactoryContainer().PostController.GetByUser)
	v1Posts.POST("", factory.FactoryContainer().PostController.Create)

	v1Users.GET("/search", factory.FactoryContainer().UserInstagramController.Search)
	v1Users.GET("/:userId/info-profile", factory.FactoryContainer().UserInstagramController.GetInfoProfile)

	v1History.GET("/get-by-user/:userId", factory.FactoryContainer().StoryController.GetByUser)

	v1Histories.POST("", factory.FactoryContainer().StoryController.Create)

	v1HistoryContents.POST("/:history_id", factory.FactoryContainer().StoryController.Create)

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
