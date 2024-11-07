package dframe

import (
	"errors"
	"math"
)

type Color struct { // A float64 RGB color value. All color values should be between 0 and 1.
	R float64
	G float64
	B float64
}

func (c Color) Invalid() bool { // Returns if the [Color] value is invalid(outside of the range 0-1, or is NaN)
	return c.R < 0 || c.G < 0 || c.B < 0 || c.R > 1 || c.G > 1 || c.B > 1 || math.IsNaN(c.R) || math.IsNaN(c.G) || math.IsNaN(c.B)
}

func Lerp(a float64, b float64, f float64) float64 { // Linearly interpolates between two float64 values.
	return a*(1.0-f) + (b * f)
}

func ColorLerp(color1, color2 Color, t float64) Color { // Linearly interpolates between two [Color] values.
	color1.R = math.Round(Lerp(float64(color1.R), float64(color2.R), t))
	color1.G = math.Round(Lerp(float64(color1.G), float64(color2.G), t))
	color1.B = math.Round(Lerp(float64(color1.B), float64(color2.B), t))
	return color1
}

var (
	InvalidColor = Color{R: math.NaN(), G: math.NaN(), B: math.NaN()} // Invalid [Color](all values = NaN)
)

var (
	ErrInvalidColor = errors.New("invalid color") // Error of a [Color] being unexpectedly invalid.
)

func FromRGB(r float64, g float64, b float64) (Color, error) { // Takes RGB values and converts them to a [Color]. If the color is invalid, returns an error and the invalid color.
	var c = Color{R: r, G: g, B: b}
	if c.Invalid() {
		return c, ErrInvalidColor
	}
	return c, nil
}

func FromRGBPanicIfErr(r float64, g float64, b float64) Color { // Same as [FromRGB], but panics if FromRGB would return an error.
	var c, err = FromRGB(r, g, b)
	if err != nil {
		panic(err)
	}
	return c
}
