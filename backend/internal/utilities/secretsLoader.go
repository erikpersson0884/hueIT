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

type HueSecrets struct {
	BaseUrl        string   `json:"hue_base_url"`
	MapDescription string   `json:"map_description"`
	LightMap []Light `json:"lightsMap"`
}

func (secrets *HueSecrets) GetLightFromMap(id uint16) (Light, error) {
	for _, light := range secrets.LightMap {
		if light.Id == id {
			return light, nil
		}
	}
	return Light{}, errors.New(fmt.Sprintf("No light with id %d", id))
}

type Light struct {
	Id uint16 `json:"id"`
	X uint `json:"x"`
	Y uint `json:"y"`
}

func LoadSecrets() (*HueSecrets, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file")
	} else {
		log.Println("Loaded environment variables from .env file")
	}

	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	secrets := HueSecrets{
		BaseUrl: os.Getenv("hue_base_url"),
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &secrets)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return &secrets, nil
}