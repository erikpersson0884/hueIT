package endpoints

import "github.com/viddem/huego/internal/utilities"

var config *utilities.HueConfig

func Init(conf  *utilities.HueConfig) {
	config = conf
}

type ErrorResponse struct {
	Message string `json:"message"`
}

