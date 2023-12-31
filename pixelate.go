package main

import (
	"image"
	"image/color"
	"image/draw"
)

func Pixelate(src image.Image, tileSize int) image.Image {
	size := src.Bounds()
	out := image.NewNRGBA(size)

	for y := size.Min.Y; y < size.Max.Y; y += tileSize {
		for x := size.Min.X; x < size.Max.X; x += tileSize {
			// Figure out color
			palette := make(color.Palette, 0, 25)
			for i := 0; i < tileSize; i++ {
				for j := 0; j < tileSize; j++ {
					current := src.At(x+j, y+i)
					palette = append(palette, current)
				}
			}
			color := Average(palette)

			// Draw tile
			rect := image.Rect(x, y, x+tileSize, y+tileSize)
			draw.Draw(out, rect,
				&image.Uniform{color}, image.ZP, draw.Src)
		}
	}

	return out
}
