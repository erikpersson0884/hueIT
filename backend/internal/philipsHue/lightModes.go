package philipsHue

import (
	"github.com/viddem/huego/internal/utilities"
	"log"
	"math/rand"
	"time"
)

const maxUInt16Val uint16 = 65535

func Neutral(config *utilities.HueConfig) {
	for _, light := range config.LightMap {
		err := SetLampCall(&utilities.LampData{
			On:         true,
			Brightness: 50,
			Hue:        30000,
			Saturation: 50,
		}, config, light.Id)

		if err != nil {
			log.Fatalf("Failed to set light mode due to err %s\n", err)
		}
	}
}

func Wave(config *utilities.HueConfig, brightness uint8) {
	var hue uint16 = 4000
	for {
		hue += 2000 % maxUInt16Val
		for _, light := range config.LightMap {
			hueData := utilities.LampData{
				On:         true,
				Brightness: brightness,
				Hue:        hue,
				Saturation: 255,
			}

			//fmt.Println(fmt.Sprintf("Should set utilities for lamp %d", light))
			err := SetLampCall(&hueData, config, light.Id)

			if err != nil {
				log.Fatalf("Failed to set light mode due to err %s\n", err)
			}

			time.Sleep(time.Second / 10)
		}
		time.Sleep(time.Second / 4)
	}
}

func randHue() uint16 {
	return uint16(rand.Int31n(int32(maxUInt16Val)))
}

func Disco(config *utilities.HueConfig) {
	for {
		for _, light := range config.LightMap {
			hueData := utilities.LampData{
				On:         true,
				Brightness: 255,
				Hue:        randHue(),
				Saturation: 255,
			}

			err := SetLampCall(&hueData, config, light.Id)

			if err != nil {
				log.Fatalf("Failed to set light mode due to err %s\n", err)
			}

			time.Sleep(time.Second / 10)
		}
		time.Sleep(time.Second / 4)
	}
}

func ChunkyDisco(config *utilities.HueConfig) {
	for {
		for _, light := range config.LightMap {
			hueData := utilities.LampData{
				On:         true,
				Brightness: 255,
				Hue:        randHue(),
				Saturation: 255,
			}

			err := SetLampCall(&hueData, config, light.Id)

			if err != nil {
				log.Fatalf("Failed to set light mode due to err %s\n", err)
			}

			time.Sleep(time.Second / 2)
		}
	}
}
