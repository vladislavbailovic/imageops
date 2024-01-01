package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

/// How many times we will attempt to recenter the
/// k-means cluster centroids in attempt to discover
/// all clusters before we give up
const MaxRecenteringAttempts int = 250

/// Once all the k-means clusters have been discovered,
/// how many additional times we will recenter the
/// centroids to have them converge to a decent value
const ConvergeanceAttempts int = 10

func Closest(point color.Color, palette color.Palette) int {
	which := -1
	maxDistance := int64(math.MaxInt64)
	for i, mean := range palette {
		dist := Distance(point, mean)
		if dist < maxDistance {
			which = i
			maxDistance = dist
		}
	}
	return which
}

func Average(cluster color.Palette) color.Color {
	sumR := 0
	sumG := 0
	sumB := 0
	for _, point := range cluster {
		r, g, b, _ := point.RGBA()
		sumR += int(r)
		sumG += int(g)
		sumB += int(b)
	}
	amount := len(cluster)
	dR := (float64(sumR/amount) / float64(math.MaxUint16)) *
		math.MaxUint8
	dG := (float64(sumG/amount) / float64(math.MaxUint16)) *
		math.MaxUint8
	dB := (float64(sumB/amount) / float64(math.MaxUint16)) *
		math.MaxUint8
	return color.RGBA{
		R: uint8(dR),
		G: uint8(dG),
		B: uint8(dB),
		A: 255,
	}
}

func ExtractFrom(img image.Image, K int) color.Palette {
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
	return Extract(samples, K)
}

/// Extract color palette using k-means clustering
func Extract(samples []color.Color, K int) color.Palette {
	rand.Seed(time.Now().UnixNano())

	// Seed centroids
	centroids := make(color.Palette, 0, K)
	for i := 0; i < K; i++ {
		centroids = append(centroids, RandomColor())
	}

	clusters := make([]color.Palette, 0, K)
	for i := 0; i < K; i++ {
		clusters = append(clusters,
			make(color.Palette, 0, len(samples)))
	}

	recenterAttempt := 0
	for convergeanceAttempt := 0; convergeanceAttempt < ConvergeanceAttempts; convergeanceAttempt++ {
		// Null out clusters
		for i := 0; i < K; i++ {
			clusters[i] = make(color.Palette, 0, len(samples))
		}

		// (Re-)Classify
		for _, point := range samples {
			which := Closest(point, centroids)
			if which >= 0 {
				clusters[which] = append(clusters[which],
					point)
			}
		}

		// Adjust centroids
		if len(centroids) != len(clusters) {
			// TODO: something very wrong happened here
			return centroids
		}
		hasMissingCentroid := false
		for i, _ := range centroids {
			if len(clusters[i]) == 0 {
				hasMissingCentroid = true
				centroids[i] = RandomColor()
			} else {
				centroids[i] = Average(clusters[i])
			}
		}

		if hasMissingCentroid {
			convergeanceAttempt -= 1
		}
		recenterAttempt += 1
		if recenterAttempt >= MaxRecenteringAttempts {
			break
		}
	}

	return centroids
}
