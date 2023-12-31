package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const RecenteringAttempts int = 10
const MaxAttempts int = 250

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

	totalIter := 0
	for iter := 0; iter < RecenteringAttempts; iter++ {
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
			log.Fatalf("Centroids vs clusters: %d vs %d",
				len(centroids), len(clusters))
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
			iter -= 1
		}
		totalIter += 1
		if totalIter >= MaxAttempts {
			break
		}
	}

	return centroids
}

func RandomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Int31n(255)),
		G: uint8(rand.Int31n(255)),
		B: uint8(rand.Int31n(255)),
		A: 255,
	}
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

func Closest(point color.Color, centroids color.Palette) int {
	which := -1
	maxDistance := int64(math.MaxInt64)
	for i, mean := range centroids {
		dist := Distance(point, mean)
		if dist < maxDistance {
			which = i
			maxDistance = dist
		}
	}
	return which
}

/// Euclidean distance (linear sRGBA)
func Distance(c1, c2 color.Color) int64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()

	dR := int64(r2) - int64(r1)
	dG := int64(g2) - int64(g1)
	dB := int64(b2) - int64(b1)

	return dR*dR + dG*dG + dB*dB
}
