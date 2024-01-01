package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	extractPaletteAndRecolorize()
	// pixelateAndRecolorize()
	// pixelateAndAsciify()
}

func pixelateAndAsciify() {
	fname := "testdata/sample2.png"
	fp, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Unable to open %s: %v", fname, err)
	}
	defer fp.Close()

	img, err := png.Decode(fp)
	if err != nil {
		log.Fatalf("Unable to decode image: %v", err)
	}

	ascii := AsciifyWith(img, 7, HtmlifyBgText())
	fmt.Printf("<pre>%s</pre>\n\n", ascii)
}

func pixelateAndRecolorize() {
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

	outFname := "simplified.png"
	outf, err := os.Create(outFname)
	if err != nil {
		log.Fatalf("Could not create %s: %v", outFname, err)
	}
	defer outf.Close()

	tmp := Pixelate(img, 25)
	size := tmp.Bounds()
	out := image.NewNRGBA(size)
	palette := color.Palette{
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		color.RGBA{R: 0, G: 255, B: 0, A: 255},
		color.RGBA{R: 0, G: 0, B: 255, A: 255},
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		color.RGBA{R: 0, G: 0, B: 0, A: 255},
	}

	width := size.Max.X - size.Min.X
	height := size.Max.Y - size.Min.Y
	for i := 0; i < width*height; i++ {
		y := i / width
		x := i % width
		point := tmp.At(x, y)
		idx := Closest(point, palette)
		if idx >= 0 {
			closest := palette[idx]
			draw.Draw(out, image.Rect(x, y, x+1, y+1),
				&image.Uniform{closest}, image.ZP, draw.Src)
		}
	}

	png.Encode(outf, out)
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

	// Initialize output surface
	rec := image.NewNRGBA(size)

	recFname := "recolorized.png"
	recf, err := os.Create(recFname)
	if err != nil {
		log.Fatalf("Could not create %s: %v", recFname, err)
	}
	defer recf.Close()

	means := ExtractFrom(img, 8)

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
