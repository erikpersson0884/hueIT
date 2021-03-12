package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/viddem/huego/internal/philipsHue"
	"log"
	"net/http"
)

func GetLamps(c *gin.Context) {
	lampInfos, err := philipsHue.GetLightsInfo(config)
	if err != nil {
		log.Printf("Failed to retrieve lights information due to err: %s\n", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Failed to retrieve lights information",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"lights": lampInfos,
	})
}
