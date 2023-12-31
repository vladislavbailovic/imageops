package main

import (
	"image/color"
	"testing"
)

func Test_Intensity(t *testing.T) {
	suite := map[uint8]color.Color{
		100: White,
		30:  Red,
		0:   Black,
	}

	for want, test := range suite {
		t.Run("intensity", func(t *testing.T) {
			got := Intensity(test)
			if want != got {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

func Test_colorToHexString(t *testing.T) {
	suite := map[string]color.Color{
		"#ffff0bff": color.RGBA{
			R: 255,
			G: 255,
			B: 11,
			A: 255,
		},
		"#ff0000ff": Red,
		"#00ff00ff": Green,
		"#0000ffff": Blue,
		"#000000ff": Black,
		"#ffffffff": White,
	}
	for want, test := range suite {
		t.Run(want, func(t *testing.T) {
			got := colorToHexString(test)
			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
