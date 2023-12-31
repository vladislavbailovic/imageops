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
