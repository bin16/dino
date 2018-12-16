package colors

import (
	"image/color"
	"math"
)

// colorsIndex

type colorsMap struct {
	p  int
	s  int
	b  int
	bg int
}

// pixPalette

type pixPalette struct {
	p  color.Color
	s  color.Color
	b  color.Color
	bg color.Color
}

func (pal *pixPalette) P() color.Color {
	return pal.p
}

func (pal *pixPalette) S() color.Color {
	return pal.s
}

func (pal *pixPalette) B() color.Color {
	return pal.b
}

func (pal *pixPalette) BG() color.Color {
	return pal.bg
}

// RGB

type RGB struct {
	R, G, B float64
}

func (rgb *RGB) toRGBA() color.RGBA {
	r := uint8(rgb.R * 255)
	g := uint8(rgb.G * 255)
	b := uint8(rgb.B * 255)

	return color.RGBA{r, g, b, 255}
}

// HSV

type HSV struct {
	H, S, V float64
}

func (hsv *HSV) HSV() (h, s, v float64) {
	return hsv.H, hsv.S, hsv.V
}

func (hsv HSV) HSVasInt() (h, s, v int) {
	return int(hsv.H), int(hsv.S), int(hsv.V)
}

func (hsv HSV) toRGBA() color.RGBA {
	rgb := hsv.toRGB()
	return rgb.toRGBA()
}

func (hsv HSV) AsRGBA() color.RGBA {
	return hsv.toRGBA()
}

func (hsv HSV) toRGB() RGB {
	var h, s, v = hsv.HSV()
	if h < 0 {
		h += 360
	}
	h /= 360
	s /= 100
	v /= 100
	var r, g, b float64

	var i = float64(math.Floor(h * 6))
	f := h*6 - i
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	switch int(i) % 6 {
	case 0:
		r = v
		g = t
		b = p
		break
	case 1:
		r = q
		g = v
		b = p
		break
	case 2:
		r = p
		g = v
		b = t
		break
	case 3:
		r = p
		g = q
		b = v
		break
	case 4:
		r = t
		g = p
		b = v
		break
	case 5:
		r = v
		g = p
		b = q
		break
	}

	return RGB{r, g, b}
}

func NewHSV(h, s, v int) HSV {
	return HSV{
		float64(h),
		float64(s),
		float64(v),
	}
}

func NewRGB(r, g, b float64) RGB {
	return RGB{r, g, b}
}
