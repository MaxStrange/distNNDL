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
	for _, l := range net.nodes {
		for _, n := range l {
			n.initialize()
		}
	}
}

// To be called AFTER a synchronous call to parseNetFromNNDL
func (net *network) train(dataChan <-chan string, mainChan chan<- string) {
	var dataIndex int
	var totalErr [][]float64
	for {
		rawData, ok := <-dataChan
		if !ok {
			break
		}
		dataIndex++

		fmt.Println(rawData)
		input, label := parseData(rawData)
		yHat := net.forward(input)

		fmt.Fprintf(os.Stdout, "yHat: %v\n", yHat)
		err := errorFunction(label, yHat)
		totalErr = append(totalErr, err)

		fmt.Fprintf(os.Stdout, "Err: %v\n", err)
		if dataIndex%net.batchLength == 0 {
			avgErr := averageError(totalErr)

			fmt.Fprintf(os.Stdout, "AvgErr: %v\n", avgErr)
			for nodeIndex, n := range net.nodes[len(net.nodes)-1] {
				e := avgErr[nodeIndex]
				fmt.Fprintf(os.Stdout, "  Loading error: %f\n", e)
				n.loadError(e, label[n.myIndex])
			}

			fmt.Fprintf(os.Stdout, "Backproping...\n")
			net.backward()

			mainChan <- net.weightsToString()
			totalErr = [][]float64{}
		}
		if dataIndex%net.mixLength == 0 {
			// TODO average the weights amongst the group
		}
	}
	mainChan <- net.weightsToString()
}

func averageError(totalErr [][]float64) []float64 {
	var avgErr []float64
	l := len(totalErr[0])
	for i := 0; i < l; i++ {
		var e float64
		for j := 0; j < len(totalErr); j++ {
			e += totalErr[j][i]
		}
		avgErr = append(avgErr, e)
	}
	return avgErr
}

func (net *network) forward(inputVector []float64) []float64 {
	for nodeIndex, n := range net.nodes[0] {
		n.loadInput(inputVector[nodeIndex])
	}
	var yHat []float64
	for _, n := range net.nodes[len(net.nodes)-1] {
		yHat = append(yHat, n.forward())
	}
	return yHat
}

func (net *network) backward() {
	for _, n := range net.nodes[0] {
		n.backward()
	}
	for _, l := range net.nodes {
		for _, n := range l {
			n.updateWeights()
			n.invalidateCache()
		}
	}
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
