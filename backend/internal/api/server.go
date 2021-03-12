package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/viddem/huego/internal/api/endpoints"
	"github.com/viddem/huego/internal/utilities"
	"log"
	"net/http"
)

var config *utilities.HueConfig

func Init(conf *utilities.HueConfig) {
	router := gin.Default()
	store := cookie.NewStore([]byte(conf.Secret))
	router.Use(sessions.Sessions("auth", store))
	endpoints.Init(conf)
	config = conf

	v1 := router.Group("/api/")
	{
		auth := v1.Group("")
		auth.Use(CheckAuth())
		{
			auth.GET("/lamps", endpoints.GetLamps)
			auth.POST("/lamps", endpoints.SetLamps)
			auth.POST("/lamps/:id", endpoints.SetLamp)
		}
		v1.POST("/auth", endpoints.Auth)
	}

	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver due to err: %s\n", err)
	}
}

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")

		if token == nil {
			InitializeAuth(c)
			return
		}
	}
}

func InitializeAuth(c *gin.Context) {
	responseType := "response_type=code"
	clientId := fmt.Sprintf("client_id=%s", config.GammaClientId)
	redirectUri := fmt.Sprintf("redirect_uri=%s", config.GammaRedirectUri)
	response := fmt.Sprintf("%s?%s&%s&%s", config.GammaAuthorizationUri, responseType, clientId, redirectUri)

	c.Header("location", response)
	c.String(http.StatusUnauthorized, response)
	c.Abort()
}