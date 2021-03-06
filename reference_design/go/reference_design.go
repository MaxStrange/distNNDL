package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Takes a slice of labels and a slice of outputs for a given forward pass
// and returns the error vector to pass back through the network for backprop.
func errorFunction(labels []FloatXX, outputs []FloatXX) (err []FloatXX) {
	for i, label := range labels {
		output := outputs[i]
		thisErr := FloatXX(math.Pow(float64(0.5*(label-output)), 2))
		err = append(err, thisErr)
	}
	return err
}

// Derivative of error function
func derError(thisNodeOutput FloatXX, thisNodeLabel FloatXX) FloatXX {
	return thisNodeOutput - thisNodeLabel
}

// Small random numbers weight initializer
func smallRandomNumbers(_layerIndex int, _inIndex int, _outIndex int) (weight FloatXX) {
	const epsilon = 0.1
	var neg float32
	if rand.Float32() > 0.5 {
		neg = 1.0
	} else {
		neg = -1.0
	}
	weight = FloatXX(rand.Float32() * epsilon * neg)
	return weight
}

// Logistic function activation
func logistic(dotProduct FloatXX) FloatXX {
	return FloatXX(1.0 / (1.0 + math.Pow(math.E, -float64(dotProduct))))
}

// Derivative of the logistic activation function
func derLogistic(z FloatXX) FloatXX {
	return z * (1.0 - z)
}

func main() {
	// Data Set
	inputVectorsAndSet := [][]FloatXX{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	labelVectorsAndSet := [][]FloatXX{{0}, {0}, {0}, {1}}

	// Some params
	weightInitializer := smallRandomNumbers
	activationFunction := logistic
	derActivationFunction := derLogistic
	learningRate := FloatXX(0.3)
	derError := derError

	inNodes := []Node{
		*NewNode(weightInitializer, activationFunction, derActivationFunction, 0, learningRate, 0, derError, InputNode),
		*NewNode(weightInitializer, activationFunction, derActivationFunction, 1, learningRate, 0, derError, InputNode),
	}
	hNodes := []Node{
		*NewNode(weightInitializer, activationFunction, derActivationFunction, 0, learningRate, 1, derError, HiddenNode),
		*NewNode(weightInitializer, activationFunction, derActivationFunction, 1, learningRate, 1, derError, HiddenNode),
	}
	outNodes := []Node{
		*NewNode(weightInitializer, activationFunction, derActivationFunction, 0, learningRate, 2, derError, OutputNode),
	}

	hNodes[0].inputNodes = []*Node{&inNodes[0], &inNodes[1]}
	hNodes[1].inputNodes = []*Node{&inNodes[0], &inNodes[1]}
	outNodes[0].inputNodes = []*Node{&hNodes[0], &hNodes[1]}

	inNodes[0].outputNodes = []*Node{&hNodes[0], &hNodes[1]}
	inNodes[1].outputNodes = []*Node{&hNodes[0], &hNodes[1]}
	hNodes[0].outputNodes = []*Node{&outNodes[0]}
	hNodes[1].outputNodes = []*Node{&outNodes[0]}

	// Layers
	in := layer{
		nodes:      inNodes,
		index:      0,
		bias:       1.0,
		biasWeight: 0.6,
	}
	hidden := layer{
		nodes:      hNodes,
		index:      1,
		bias:       1.0,
		biasWeight: 0.35,
	}
	out := layer{
		nodes:      outNodes,
		index:      2,
		bias:       0.0,
		biasWeight: 0.0,
	}

	// Network
	net := network{
		inputLayer:   in,
		hiddenLayers: []layer{hidden},
		outputLayer:  out,
		errFunc:      errorFunction,
	}

	// Print the crap so it doesn't harass me about things being unused
	fmt.Println(fmt.Sprintf("Input vectors: %v", inputVectorsAndSet))
	fmt.Println(fmt.Sprintf("Labels: %v", labelVectorsAndSet))
	fmt.Println(net.String())

	net.Initialize()
	net.PrintWeights()

	//net.Run(inputVectorsAndSet, labelVectorsAndSet, 10)
}
