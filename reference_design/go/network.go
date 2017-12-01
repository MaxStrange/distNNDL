package main

import (
	"bytes"
	"fmt"
	"log"
)

type network struct {
	inputLayer   layer
	hiddenLayers []layer
	outputLayer  layer
	errFunc      func(labels []FloatXX, outputs []FloatXX) (errors []FloatXX)
}

func (net *network) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Network of sizes %d -> ", len(net.inputLayer.nodes)))
	for _, lay := range net.hiddenLayers {
		buffer.WriteString(fmt.Sprintf("%d -> ", len(lay.nodes)))
	}
	buffer.WriteString(fmt.Sprintf("%d", len(net.outputLayer.nodes)))
	return buffer.String()
}

func (net *network) Initialize() {
	for _, node := range net.AllNeurons() {
		node.Initialize()
	}
}

// Run runs the training for the given number of epochs.
func (net *network) Run(inputVectors [][]FloatXX, labels [][]FloatXX, numEpochs int) {
	if len(inputVectors) != len(labels) {
		log.Fatalf("Len(inputVectors) != len(labels)")
	} else if len(inputVectors) <= 0 {
		log.Fatalf("Len(inputVectors) <= 0")
	} else if len(labels[0]) != len(net.outputLayer.nodes) {
		log.Fatalf("Label size must be equal to outputLayer size")
	} else if len(inputVectors[0]) != len(net.inputLayer.nodes) {
		log.Fatalf("Input vector size must be equal to input layer size")
	}

	var accuracies []FloatXX

	for epochIndex := 0; epochIndex < numEpochs; epochIndex++ {
		fmt.Println("-------------------------------")
		fmt.Println(fmt.Sprintf("EPOCH %d:", epochIndex))
		fmt.Println("-------------------------------")
		for vecAndLabelIndex, vec := range inputVectors {
			label := labels[vecAndLabelIndex]
			fmt.Println(fmt.Sprintf("DATA: %v, LABEL: %v", vec, label))
			for nodeIndex, node := range net.inputLayer.nodes {
				node.LoadInput(vec[nodeIndex])
			}
			var yHat []FloatXX
			for _, node := range net.outputLayer.nodes {
				yHat = append(yHat, node.Forward())
			}
			err := net.errFunc(label, yHat)
			for nodeIndex, node := range net.outputLayer.nodes {
				e := err[nodeIndex]
				node.LoadError(e, label[node.myIndex])
			}
			for _, node := range net.inputLayer.nodes {
				node.Backward()
			}
			for _, node := range net.AllNeurons() {
				node.UpdateWeights()
				node.InvalidateCache()
			}
		}
		if epochIndex%10 == 0 {
			accuracies = append(accuracies, net.Evaluate(inputVectors, labels))
		}
	}
	// TODO: Plot
}

// Evaluate evaluates the accuracy of the network on the given inputVectors.
func (net *network) Evaluate(inputVectors [][]FloatXX, labels [][]FloatXX) FloatXX {
	totalAcc := FloatXX(0.0)
	for vecIndex, vec := range inputVectors {
		label := labels[vecIndex]
		for nodeIndex, node := range net.inputLayer.nodes {
			inputVal := vec[nodeIndex]
			node.LoadInput(inputVal)
		}
		var yHat []FloatXX
		for _, node := range net.outputLayer.nodes {
			yHat = append(yHat, node.Forward())
		}
		acc := int(yHat[0]+0.5) == int(label[0]+0.5)
		if acc {
			totalAcc++
		}
		for _, node := range net.AllNeurons() {
			node.InvalidateCache()
		}
	}
	totalAcc /= FloatXX(len(inputVectors))
	return totalAcc
}

func (net *network) AllNeurons() []Node {
	var allNodes []Node
	allNodes = append(allNodes, net.inputLayer.nodes...)
	for _, hLayer := range net.hiddenLayers {
		allNodes = append(allNodes, hLayer.nodes...)
	}
	allNodes = append(allNodes, net.outputLayer.nodes...)
	return allNodes
}
