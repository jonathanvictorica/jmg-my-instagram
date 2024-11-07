package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"jmg-my-instagram-ms/config"
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/routes"
	"log"
)

func main() {

	configuration, err := config.LoadConfig("./config") // "." indica el directorio actual
	if err != nil {
		log.Fatalf("No se pudo cargar la configuración: %v", err)
	}

	_, err = config.OpenConnection(configuration)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	err = repository.AutoMigrate()
	if err != nil {
		log.Fatalf("No se pudo crear la base de datos: %v", err)
	}

	router := gin.Default()
	routes.SetupRouter(router)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	err = router.Run(":" + config.Configuration.Server.Port)
	if err != nil {
		return
	}

}
