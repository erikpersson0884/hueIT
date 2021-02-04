package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viddem/huego/internal/philipsHue"
	"github.com/viddem/huego/internal/utilities"
	"log"
	"net/http"
)

var secrets *utilities.HueSecrets

func main() {
	sec, err := utilities.LoadSecrets()
	if err != nil {
		log.Fatalf("Failed to load secrets due to err: %s\n", err)
	}
	secrets = sec

	go philipsHue.Wave(sec, 120)
	//err = philipsHue.Disco(sec)
	//err = philipsHue.ChunkyDisco(sec)
	//err = philipsHue.Neutral(sec)

	router := gin.Default()

	v1 := router.Group("/api/")
	{
		v1.GET("/lamps", getLamps)
		v1.POST("/lamps", setLamps)
		v1.POST("/lamp/:id", setLamp)
	}
	err = router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver due to err: %s\n", err)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func getLamps(c *gin.Context) {
	lampInfos, err := philipsHue.GetLightsInfo(secrets)
	if err != nil {
		log.Printf("Failed to retrieve lights information due to err: %s\n", err)
		c.JSON(500, ErrorResponse{
			Message: "Failed to retrieve lights information",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"lights": lampInfos,
	})
}

func setLamps(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": http.StatusInternalServerError,
		"error":  "Not implemented",
	})
}

func setLamp(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": http.StatusInternalServerError,
		"error":  "Not implemented",
	})
}
