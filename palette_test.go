package main

import (
	"image/color"
	"testing"
)

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

func Test_Extract(t *testing.T) {
	test := []color.Color{
		Red, Red, Red,
		Green, Green,
		Blue,
	}
	palette := Extract(test, 3)
	if len(palette) != 3 {
		t.Errorf("expected 3 color palette, got %d",
			len(palette))
	}

	for _, val := range test {
		hasColor := false
		for _, c := range palette {
			if c == val {
				hasColor = true
				break
			}
		}
		if !hasColor {
			t.Errorf("did not find %v in palette %v",
				val, palette)
		}
	}
}
