package main

import (
	"image/color"
	"testing"
)

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

func Test_ToHexString(t *testing.T) {
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
			got := ToHexString(test)
			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
