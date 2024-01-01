package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	// extractPaletteAndRecolorize()
	pixelateAndRecolorize()
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

	palette := ExtractFrom(img, 8)

	png.Encode(outf, Recolor(Pixelate(img, 25), palette))
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

	recFname := "recolorized.png"
	recf, err := os.Create(recFname)
	if err != nil {
		log.Fatalf("Could not create %s: %v", recFname, err)
	}
	defer recf.Close()

	palette := ExtractFrom(img, 8)
	rec := Recolor(img, palette)

	png.Encode(recf, rec)
}
