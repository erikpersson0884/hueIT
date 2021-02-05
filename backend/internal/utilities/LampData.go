package utilities

import "math"

type LampData struct {
	On         bool   `json:"on"`
	Brightness uint8  `json:"bri"`
	Hue        uint16 `json:"hue"`
	Saturation uint8  `json:"sat"`
}

const maxUint8Val = 255

func (l *LampData)GetHSB() HSB {
	return HSB{
		H: float64(l.Hue) / maxUint8Val,
		S: float64(l.Saturation) / maxUint8Val,
		B: float64(l.Brightness) / maxUint8Val,
	}
}

func (l *LampData)ToRGB() LampDataRGB {
	return LampDataRGB{
		On:  l.On,
		Rgb: hsbToRgb(l.GetHSB()),
	}
}

type LampDataRGB struct {
	On bool `json:"on"`
	Rgb RGB `json:"rgb"`
}


type SimpleLampData struct {
	On bool `json:"on"`
	Hue float64 `json:"h"`
	Saturation float64 `json:"s"`
	Brightness float64 `json:"b"`
}

const EightMaxVal = 254
const HueMaxVal = 65535

func (l *LampData) Simplify() SimpleLampData {
	return SimpleLampData{
		On:         l.On,
		Hue:        roundFloat((float64(l.Hue) / HueMaxVal) * 360),
		Saturation: roundFloat(float64(l.Saturation) / EightMaxVal * 100),
		Brightness: roundFloat(float64(l.Brightness) / EightMaxVal * 100),
	}
}

func roundFloat(a float64) float64 {
	return math.Round(a * 100) / 100
}