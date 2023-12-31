package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
	AsciiRangeStart int = 32  // "space"
	AsciiRangeEnd   int = 126 // "tilda"
)

type AsciiPalette map[uint8]string

func getAsciiPalette() AsciiPalette {
	point := fixed.Point26_6{fixed.I(7), fixed.I(13)}
	bounds := image.Rect(0, 0, 28, 28)

	minimum := math.MaxInt
	maximum := 0
	byLetter := make(map[string]int,
		AsciiRangeEnd-AsciiRangeStart)
	for idx := AsciiRangeStart; idx < AsciiRangeEnd; idx++ {
		img := image.NewNRGBA(bounds)
		render := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(White),
			Face: basicfont.Face7x13,
			Dot:  point,
		}
		draw.Draw(render.Dst, bounds,
			image.NewUniform(Black), image.ZP, draw.Src)
		render.DrawString(string(rune(idx)))
		area := (bounds.Max.X - bounds.Min.X) *
			(bounds.Max.Y - bounds.Min.Y)
		intensity := 0
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, _ := render.Dst.At(x, y).RGBA()
				intensity += int(r + g + b)
			}
		}
		avgIntensity := intensity / area

		if minimum > avgIntensity && avgIntensity > 0 {
			minimum = avgIntensity
		}
		if maximum < avgIntensity {
			maximum = avgIntensity
		}
		byLetter[string(rune(idx))] = avgIntensity
	}

	ranges := make(map[uint8]string, len(byLetter))
	for letter, intensity := range byLetter {
		if intensity != 0 {
			intensity -= minimum
		}
		amount := float64(intensity) / float64(maximum)
		gray := uint8(amount*10) * 10
		ranges[gray] = letter
	}
	ranges[0] = " "

	return ranges
}

func Intensity(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	dR := float64(r) / float64(math.MaxUint16)
	dG := float64(g) / float64(math.MaxUint16)
	dB := float64(b) / float64(math.MaxUint16)
	amount := (dR + dG + dB) / 3
	return uint8(amount*10) * 10
}

// TODO: very similar to Pixelate
func Asciify(src image.Image, tileSize int) string {
	size := src.Bounds()
	width := size.Max.X - size.Min.X
	height := size.Max.Y - size.Min.Y
	out := ""

	ascii := getAsciiPalette()

	for y := size.Min.Y; y < size.Max.Y; y += tileSize {
		for x := size.Min.X; x < size.Max.X; x += tileSize {
			// Figure out color
			palette := make(color.Palette, 0, width*height)
			for i := 0; i < tileSize; i++ {
				for j := 0; j < tileSize; j++ {
					current := src.At(x+j, y+i)
					palette = append(palette, current)
				}
			}
			color := Average(palette)
			// Asciify color, step 1: map to shade intensity
			intensity := Intensity(color)
			// Asciify color, step 2: map shade to character
			out += ascii[intensity]
		}
		out += "\n"
	}

	return out
}
