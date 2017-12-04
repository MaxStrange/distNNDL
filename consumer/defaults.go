// This file provides a bunch of default functions for activation/derivatives, etc.
package main

import (
	"math"
	"math/rand"
)

func derLogistic(z float64) float64 {
	return z * (1.0 - z)
}

func logistic(dotProduct float64) float64 {
	return float64(1.0 / (1.0 + math.Pow(math.E, -float64(dotProduct))))
}

func smallRandomNumbers(_layerIndex int, _intIndex int, _outIndex int) (weight float64) {
	const epsilon = 0.1
	var neg float64
	if rand.Float64() > 0.5 {
		neg = 1.0
	} else {
		neg = -1.0
	}
	weight = rand.Float64() * epsilon * neg
	return weight
}

func derError(thisNodeOutput float64, thisNodeLabel float64) float64 {
	return thisNodeOutput - thisNodeLabel
}

func errorFunction(labels []float64, outputs []float64) (err []float64) {
	for i, label := range labels {
		output := outputs[i]
		thisErr := math.Pow(0.5*(label-output), 2)
		err = append(err, thisErr)
	}
	return err
}
