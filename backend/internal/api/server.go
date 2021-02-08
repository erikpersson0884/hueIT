package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viddem/huego/internal/api/endpoints"
	"github.com/viddem/huego/internal/utilities"
	"log"
)

func Init(secrets *utilities.HueSecrets) {
	router := gin.Default()
	endpoints.Init(secrets)

	v1 := router.Group("/api/")
	{
		v1.GET("/lamps", endpoints.GetLamps)
		v1.POST("/lamps", endpoints.SetLamps)
		v1.POST("/lamps/:id", endpoints.SetLamp)
	}
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver due to err: %s\n", err)
	}
}
