package main

import (
	"image/color"
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
			iter -= 1
		}
		totalIter += 1
		if totalIter >= MaxAttempts {
			break
		}
	}

	return centroids
}
