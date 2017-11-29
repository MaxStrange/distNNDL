package main

import (
	"bytes"
	"fmt"
)

type network struct {
	inputLayer   layer
	hiddenLayers []layer
	outputLayer  layer
	errFunc      func(labels []floatXX, outputs []floatXX) (errors []floatXX)
}

func (net network) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Network of sizes %d -> ", len(net.inputLayer.nodes)))
	for _, lay := range net.hiddenLayers {
		buffer.WriteString(fmt.Sprintf("%d -> ", len(lay.nodes)))
	}
	buffer.WriteString(fmt.Sprintf("%d", len(net.outputLayer.nodes)))
	return buffer.String()
}
