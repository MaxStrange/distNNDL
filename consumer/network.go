package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type network struct {
	batchLength int
	epochLength int
	mixLength   int
	nodes       [][]*node // nodes[layerIndex][nodeIndex]
	weights     []float64
}

func makeNetwork(batchLength int, epochLength int, mixLength int) *network {
	return &network{
		batchLength: batchLength,
		epochLength: epochLength,
		mixLength:   mixLength,
		nodes:       [][]*node{},
		weights:     []float64{},
	}
}

func (net *network) parseNetFromNNDL(nndl string) {
	parseNNDLIntoNetwork(nndl, net)
	fmt.Println("Parsed network from input. Here it is:")
	fmt.Println(net.string())
	fmt.Println("Initializing network weights")
	for _, l := range net.nodes {
		for _, n := range l {
			n.initialize()
		}
	}
	fmt.Fprintf(os.Stdout, "After initialization: %s\n", net.weightsToString())
}

// To be called AFTER a synchronous call to parseNetFromNNDL
func (net *network) train(dataChan <-chan string, mainChan chan<- string) {
	fmt.Fprintf(os.Stdout, "Before training, weights are: %s\n", net.weightsToString())
	var dataIndex int
	for {
		rawData, ok := <-dataChan
		if !ok {
			break
		}
		dataIndex++
		fmt.Println(rawData)

		input, label := parseData(rawData)
		net.forward(input)
		if dataIndex%net.batchLength == 0 {
			net.backward(label)
			mainChan <- net.weightsToString()
		}
		if dataIndex%net.mixLength == 0 {
			// TODO average the weights amongst the group
		}
	}
	mainChan <- net.weightsToString()
}

func (net *network) forward(inputVector []float64) {
	for nodeIndex, n := range net.nodes[0] {
		n.loadInput(inputVector[nodeIndex])
	}
	var yHat []float64
	for _, n := range net.nodes[len(net.nodes)-1] {
		yHat = append(yHat, n.forward())
	}
	fmt.Fprintf(os.Stdout, "y_hat: %v\n", yHat)
}

func (net *network) backward(labelVector []float64) {
	// TODO
}

func parseData(rawData string) (input []float64, label []float64) {
	splitStr := strings.Split(rawData, ";")
	instr := splitStr[0]
	labelStr := splitStr[1]
	for _, number := range strings.Split(instr, ", ") {
		floatNum, err := strconv.ParseFloat(number, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Got non-integer input value: %s\n", number)
			os.Exit(-1)
		}
		input = append(input, floatNum)
	}
	for _, number := range strings.Split(labelStr, ", ") {
		floatNum, err := strconv.ParseFloat(number, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Got non-integer label value: %s\n", number)
			os.Exit(-1)
		}
		label = append(label, floatNum)
	}
	return input, label
}

func (net *network) weightsToString() string {
	var buffer bytes.Buffer
	for _, nodeList := range net.nodes {
		for _, n := range nodeList {
			buffer.WriteString(fmt.Sprintf("%s: ", n.id()))
			for _, w := range n.weights {
				buffer.WriteString(fmt.Sprintf("%f ", w))
			}
		}
	}
	return buffer.String()
}

func (net *network) string() string {
	var buffer bytes.Buffer
	for layerIndex, layer := range net.nodes {
		buffer.WriteString(fmt.Sprintf("Layer %d\n", layerIndex))
		for _, n := range layer {
			buffer.WriteString(fmt.Sprintf("  %d -> [", n.myIndex))
			for _, otherNode := range n.outputNodes {
				buffer.WriteString(fmt.Sprintf(" %d_%d ", otherNode.layerIndex, otherNode.myIndex))
			}
			buffer.WriteString(fmt.Sprintf("]\n"))
		}
	}
	return buffer.String()
}
