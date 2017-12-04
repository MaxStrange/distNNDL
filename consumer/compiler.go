package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type simpleNet struct {
	connections [][]string
	layers      [][]string
	layerNames  []string
}

func parseNNDLIntoNetwork(nndlContent string, net *network) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	pypath := path.Join(exPath, "python", "compiler.py")
	pycmd := exec.Command("python3", pypath)
	pyIn, err := pycmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	pyOut, err := pycmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	pyErr, err := pycmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	pycmd.Start()
	pyIn.Write([]byte(nndlContent))
	pyIn.Close()
	pyBytes, err := ioutil.ReadAll(pyOut)
	if err != nil {
		panic(err)
	}
	pyErrBytes, err := ioutil.ReadAll(pyErr)
	if err != nil {
		panic(err)
	}
	pycmd.Wait()
	if string(pyErrBytes) != "" {
		fmt.Println(string(pyErrBytes))
		fmt.Fprintf(os.Stderr, "Error received from python subprocess.")
		os.Exit(-1)
	}

	var netFromPy map[string][][]string
	err = json.Unmarshal(pyBytes, &netFromPy)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	simple := simpleNet{
		connections: netFromPy["connections"],
		layers:      netFromPy["layers"],
		layerNames:  netFromPy["layer_names"][0],
	}

	const learningRate = 0.3
	for layerIndex, layer := range simple.layers {
		var nt nodeType
		if layerIndex == 0 {
			nt = inputNode
		} else if layerIndex == len(simple.layers)-1 {
			nt = outputNode
		} else {
			nt = hiddenNode
		}
		constructingLayer := []*node{}
		for nodeIndex := range layer {
			newOne := newNode(
				smallRandomNumbers,
				logistic,
				derLogistic,
				nodeIndex,
				learningRate,
				layerIndex,
				derError,
				nt,
			)
			constructingLayer = append(constructingLayer, newOne)
		}
		net.nodes = append(net.nodes, constructingLayer)
	}

	for _, fromTo := range simple.connections {
		from := fromTo[0]
		to := fromTo[1]
		fromSplit := strings.Split(from, "_")
		toSplit := strings.Split(to, "_")
		fromLayerName := strings.Join(fromSplit[0:len(fromSplit)-1], "_")
		toLayerName := strings.Join(toSplit[0:len(toSplit)-1], "_")
		fromNodeIndex, _ := strconv.ParseInt(fromSplit[len(fromSplit)-1], 10, 64)
		toNodeIndex, _ := strconv.ParseInt(toSplit[len(toSplit)-1], 10, 64)

		fromLayerIndex := getIndexFromLayerName(fromLayerName, simple.layerNames)
		toLayerIndex := getIndexFromLayerName(toLayerName, simple.layerNames)
		toNode := net.nodes[toLayerIndex][toNodeIndex]
		fromNode := net.nodes[fromLayerIndex][fromNodeIndex]

		toNode.inputNodes = append(toNode.inputNodes, fromNode)
		fromNode.outputNodes = append(fromNode.outputNodes, toNode)
	}
}

func getIndexFromLayerName(layerName string, layers []string) int {
	for index, name := range layers {
		if layerName == name {
			return index
		}
	}
	return -1
}
