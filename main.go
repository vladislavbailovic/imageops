package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	extractPaletteAndRecolorize()
}

func extractPaletteAndRecolorize() {
	fname := "testdata/sample3.png"
	fp, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Unable to open %s: %v", fname, err)
	}
	defer fp.Close()

	img, err := png.Decode(fp)
	if err != nil {
		log.Fatalf("Unable to decode image: %v", err)
	}

	size := img.Bounds()

	width := size.Max.X - size.Min.X
	height := size.Max.Y - size.Min.Y
	samples := make([]color.Color, 0, width*height)
	for i := 0; i < width*height; i++ {
		y := i / width
		x := i % width
		point := img.At(x, y)
		samples = append(samples, point)
	}

	// Initialize output surface
	rec := image.NewNRGBA(size)

	recFname := "recolorized.png"
	recf, err := os.Create(recFname)
	if err != nil {
		log.Fatalf("Could not create %s: %v", recFname, err)
	}
	defer recf.Close()

	means := Extract(samples, 8)

	for i := 0; i < width*height; i++ {
		y := i / width
		x := i % width
		point := img.At(x, y)
		which := Closest(point, means)
		if which >= 0 {
			draw.Draw(rec, image.Rect(x, y, x+1, y+1), &image.Uniform{means[which]}, image.ZP, draw.Src)
		}
	}
	png.Encode(recf, rec)
}
