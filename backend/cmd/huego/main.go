package main

import (
	"github.com/viddem/huego/internal/api"
	"github.com/viddem/huego/internal/philipsHue"
	"github.com/viddem/huego/internal/utilities"
	"log"
)

func main() {
	sec, err := utilities.LoadConfigs()
	if err != nil {
		log.Fatalf("Failed to load secrets due to err: %s\n", err)
	}
	api.Init(sec)

	//go philipsHue.Wave(sec, 200)
	go philipsHue.Disco(sec)
	//go philipsHue.ChunkyDisco(sec)
	//err = philipsHue.Neutral(sec)
}
