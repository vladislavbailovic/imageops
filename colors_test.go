package main

import (
	"image/color"
	"testing"
)

var Red color.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var Green color.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var Blue color.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var Black color.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var White color.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func Test_Distance(t *testing.T) {
	want := Distance(White, Black)
	got := Distance(Black, White)
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_RandomColor(t *testing.T) {
	v1 := RandomColor()
	v2 := RandomColor()
	if v1 == v2 {
		t.Errorf("expected random colors to differ: %v", v1)
	}
}

func Test_Average(t *testing.T) {
	test := color.Palette{Red, Green}
	want := color.RGBA{R: 127, G: 127, B: 0, A: 255}
	got := Average(test)
	if want != got {
		t.Errorf("average error: want %v, got %v", want, got)
	}
}

func Test_Closest(t *testing.T) {
	palette := color.Palette{Red, Green, Blue, White, Black}
	suite := []color.Color{
		color.RGBA{R: 200, G: 0, B: 0, A: 255},
		color.RGBA{R: 0, G: 200, B: 0, A: 255},
		color.RGBA{R: 0, G: 0, B: 200, A: 255},
		color.RGBA{R: 200, G: 200, B: 200, A: 255},
		color.RGBA{R: 0, G: 0, B: 0, A: 255},
	}

	for want, test := range suite {
		t.Run("closest", func(t *testing.T) {
			got := Closest(test, palette)
			if want != got {
				t.Errorf("%v error: want %d, got %d",
					test, want, got)
			}
		})
	}
}
