package endpoints

import "github.com/viddem/huego/internal/utilities"

var secrets *utilities.HueSecrets

func Init(s *utilities.HueSecrets) {
	secrets = s
}

type ErrorResponse struct {
	Message string `json:"message"`
}

