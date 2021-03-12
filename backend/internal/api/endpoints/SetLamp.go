package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/viddem/huego/internal/philipsHue"
	"github.com/viddem/huego/internal/utilities"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type BasicLampData struct {
	On bool `json:"on"`
	Hsb utilities.HSB `json:"hsb"`
}


func (basic *BasicLampData) ToLampData() *utilities.LampData {
	b := uint8((basic.Hsb.B / 100) * utilities.EightMaxVal)
	h := uint16((basic.Hsb.H / 360) * utilities.HueMaxVal)
	s := uint8((basic.Hsb.S / 100) * utilities.EightMaxVal)

	return &utilities.LampData{
		On:         basic.On,
		Brightness: b,
		Hue:        h,
		Saturation: s,
	}
}

func SetLamp(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message:   fmt.Sprintf("Invalid lamp id %s", idString),
		})
		return
	}

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

	err = philipsHue.SetLampCall(receivedLampData.ToLampData(), config, uint16(id))
	if err != nil {
		log.Printf("Error: Failed to set lamp, err: %s\n", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to set lamp"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
