package utilities

import "math"

type RGB struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

type
RGBf struct {
	R float64
	G float64
	B float64
}

type HSB struct {
	H float64
	S float64
	B float64
}

func rgbIntToRgbFloat(rgbInt RGB) RGBf {
	return RGBf{
		R: float64(rgbInt.R) / 255.0,
		G: float64(rgbInt.G) / 255.0,
		B: float64(rgbInt.B) / 255.0,
	}
}

func rgbFloatToRgbInt(rgbF RGBf) RGB {
	return RGB{
		R: uint8(rgbF.R * 255.0),
		G: uint8(rgbF.G * 255.0),
		B: uint8(rgbF.B * 255.0),
	}
}

const threshold = 1e-6

func compareFloats(a, b float64) bool {
	if (a - b) <= threshold {
		return true
	}
	return false
}

func rgbToHsb(rgbInt RGB) HSB {
	rgb := rgbIntToRgbFloat(rgbInt)
	min := math.Min(rgb.R, math.Min(rgb.G, rgb.B))
	max := math.Max(rgb.R, math.Max(rgb.G, rgb.B))
	brightness := max

	delta := max - min

	saturation := 0.0
	hue := 0.0
	if delta < threshold {
		return HSB{
			H: hue,
			S: saturation,
			B: brightness,
		}
	}

	if max > 0 {
		saturation = delta / max
	} else {
		// Hue should be NaN but I fear that might break things and
		// online converters appear to be using a philipsHue of 0 degrees for black.
		return HSB{
			H: hue,
			S: saturation,
			B: brightness,
		}
	}
	if compareFloats(rgb.R, max) {
		hue = (rgb.G - rgb.B) / delta
	} else if compareFloats(rgb.G, max) {
		hue = 2.0 + (rgb.B-rgb.R)/delta
	} else {
		hue = 4.0 + (rgb.R-rgb.G)/delta
	}

	hue *= 60.0

	if hue <= 0 {
		hue += 360
	}

	return HSB{
		H: hue,
		S: saturation,
		B: brightness,
	}
}

func hsbToRgb(hsb HSB) RGB {
	if compareFloats(hsb.S, 0) {
		return rgbFloatToRgbInt(RGBf{
			R: hsb.B,
			G: hsb.B,
			B: hsb.B,
		})
	}

	hh := hsb.H
	if hh >= 360 {
		hh = 0.0
	}
	hh /= 60.0
	i := int64(hh)
	ff := hh - float64(i)
	p := hsb.B * (1.0 - hsb.S)
	q := hsb.B * (1.0 - (hsb.S * ff))
	t := hsb.B * (1.0 - (hsb.S * (1.0 - ff)))

	var r, g, b float64

	switch i {
	case 0:
		r = hsb.B
		g = t
		b = p
		break
	case 1:
		r = q
		g = hsb.B
		b = p
		break
	case 2:
		r = p
		g = hsb.B
		b = t
		break
	case 3:
		r = p
		g = q
		b = hsb.B
		break
	case 4:
		r = t
		g = p
		b = hsb.B
		break
	case 5:
	default:
		r = hsb.B
		g = p
		b = q
		break
	}

	return rgbFloatToRgbInt(RGBf{
		R: r,
		G: g,
		B: b,
	})
}
