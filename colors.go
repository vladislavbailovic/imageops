package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
)

var Red color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var Green color.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var Blue color.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var Black color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var White color.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func RandomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Int31n(255)),
		G: uint8(rand.Int31n(255)),
		B: uint8(rand.Int31n(255)),
		A: 255,
	}
}

func Inverse(src color.Color) color.Color {
	r, g, b, _ := src.RGBA()
	dR := (float64(r) / float64(math.MaxUint16)) *
		math.MaxUint8
	dG := (float64(g) / float64(math.MaxUint16)) *
		math.MaxUint8
	dB := (float64(b) / float64(math.MaxUint16)) *
		math.MaxUint8
	return color.RGBA{
		R: 255 - uint8(dR),
		G: 255 - uint8(dG),
		B: 255 - uint8(dB),
		A: 255,
	}
}

/// Euclidean distance (linear sRGBA)
func Distance(c1, c2 color.Color) int64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	dR := int64(r2) - int64(r1)
	dG := int64(g2) - int64(g1)
	dB := int64(b2) - int64(b1)

	return dR*dR + dG*dG + dB*dB
}

func ToHexString(c color.Color) string {
	r, g, b, a := c.RGBA()
	res := 0 |
		((r & 0xff) << 24) |
		((g & 0xff) << 16) |
		((b & 0xff) << 8) |
		(a & 0xff)
	return fmt.Sprintf("#%08x", res)
}
