package endpoints

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type authRequest struct {
	Code *string `json:"code"`
}

type gammaTokenResponse struct {
	AccessToken string `json:"code" binding:"required"`
	ExpiresIn int64 `json:"expires_in" binding:"required"`
}

func Auth(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error: failed to read json data, err: %s\n", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Failed to read request data",
		})
		return
	}

	var receivedCode authRequest
	err = json.Unmarshal(jsonData, &receivedCode)
	if receivedCode.Code == nil {
		log.Printf("No code in request")
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Invalid or missing code",
		})
		return
	}

	authVal := fmt.Sprintf("%s:%s", config.GammaClientId, config.GammaSecret)
	b64EncodedAuth := base64.StdEncoding.EncodeToString([]byte(authVal))

	url := fmt.Sprintf("%s?grant_type=authorization_code&client_id=%s&redirect_uri=%s&code=%s", config.GammaTokenUri, config.GammaClientId, config.GammaRedirectUri, *receivedCode.Code)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", b64EncodedAuth))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		genericError(err, c)
		return
	}

	if res.StatusCode == 200 {
		var resp gammaTokenResponse
		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			genericError(err, c)
			return
		}

		session := sessions.Default(c)
		session.Set("token", resp.AccessToken)
		session.Options(sessions.Options{
			MaxAge: int(resp.ExpiresIn),
		})
		err = session.Save()
		if err != nil {
			log.Printf("Failed to create session: %v\n", err)
			c.JSON(500, ErrorResponse{
				Message: "Failed to create session",
			})
		}

		c.String(200, "Session created")
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		genericError(err, c)
		return
	}

	log.Printf("Gamma responded with %d | %s\n", res.StatusCode, string(resBody))
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: "Incorrect token",
	})
}

func genericError(err error, c *gin.Context) {
	log.Printf("Got error %s\n", err)
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Message: "Something went wrong",
	})
}