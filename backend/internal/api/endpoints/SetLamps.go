package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/viddem/huego/internal/philipsHue"
	"io/ioutil"
	"log"
	"net/http"
)

func SetLamps(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error: failed to read json data, err: %s\n", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Failed to read request data",
		})
		return
	}

	var receivedLampData BasicLampData
	err = json.Unmarshal(jsonData, &receivedLampData)
	if err != nil {
		log.Printf("Error: invalid json data, err: %s\n", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Invalid json data",
		})
		return
	}

	fmt.Printf("Received lamp data: %+v\n resulting in: %+v\n", receivedLampData, receivedLampData.ToLampData())

	err = philipsHue.SetAllLampsCall(receivedLampData.ToLampData(), config)
	if err != nil {
		log.Printf("Error: Failed to set lamp, err: %s\n", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to set lamp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}