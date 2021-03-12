package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

type HueConfig struct {
	BaseUrl               string
	MapDescription        string  `json:"map_description"`
	LightMap              []Light `json:"lightsMap"`
	GammaAuthorizationUri string
	GammaRedirectUri      string
	GammaTokenUri         string
	GammaMeUri            string
	GammaSecret           string
	GammaClientId         string
	Secret 				  string
}

func (config *HueConfig) GetLightFromMap(id uint16) (Light, error) {
	for _, light := range config.LightMap {
		if light.Id == id {
			return light, nil
		}
	}
	return Light{}, errors.New(fmt.Sprintf("No light with id %d", id))
}

type Light struct {
	Id uint16 `json:"id"`
	X  uint   `json:"x"`
	Y  uint   `json:"y"`
}

func LoadConfigs() (*HueConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file, using existing environment variables")
	} else {
		log.Println("Loaded environment variables from .env file")
	}

	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	config := HueConfig{
		BaseUrl:               loadNonEmptyString("HUE_BASE_URL"),
		GammaAuthorizationUri: loadNonEmptyString("GAMMA_AUTHORIZATION_URI"),
		GammaRedirectUri:      loadNonEmptyString("GAMMA_REDIRECT_URI"),
		GammaTokenUri:         loadNonEmptyString("GAMMA_TOKEN_URI"),
		GammaMeUri:            loadNonEmptyString("GAMMA_ME_URI"),
		GammaSecret:           loadNonEmptyString("GAMMA_SECRET"),
		GammaClientId:         loadNonEmptyString("GAMMA_CLIENT_ID"),
		Secret:				   loadNonEmptyString("SECRET"),
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	err = jsonFile.Close()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func loadNonEmptyString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable '%s' is not set or set to empty which is not allowed!", key)
	}
	return val
}