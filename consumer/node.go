package main

import "fmt"
import "os"

type nodeType int

const (
	inputNode nodeType = iota
	hiddenNode
	outputNode
)

// weightInitializer is a typedef for functions passed around for initializing weights.
type weightInitializer func(layerIndex int, inIndex int, outIndex int) (weight float64)

// activationFunction is a typedef for activation functions.
type activationFunction func(dotProduct float64) float64

// derivativeActivation is a typedef of the derivative of an activation function.
type derivativeActivation func(z float64) float64

// derivativeError is a typedef of the derivative of an error function.
type derivativeError func(thisNodeOutput float64, thisNodeLabel float64) float64

// node is the main node object in a network.
type node struct {
	inputNodes           []*node                                                          // The list of nodes that input into this one
	outputNodes          []*node                                                          // The list of nodes that this one input into
	weightInitializer    func(layerIndex int, inIndex int, outIndex int) (weight float64) // The function used to initialize the weights into this node
	activation           func(dotProduct float64) float64                                 // The activation function for this node
	derivativeActivation func(z float64) float64                                          // The derivative of the activation function
	myIndex              int                                                              // The index of this node within a layer, 0-based
	learningRate         float64                                                          // The learning rate for this node's input weights
	biasValue            float64                                                          // The value of this layer's bias node
	biasWeight           float64                                                          // The initial weight of the bias node
	layerIndex           int                                                              // The index of the layer that this node is a part of, 0-based
	derivativeError      func(thisNodeOutput float64, thisNodeLabel float64) float64      // The derivative of the error function
	weights              []float64                                                        // The weights, lined up with inputNodes
	myType               nodeType                                                         // The type of this node
	inputValue           float64                                                          // The input value into this node - only used by Input Nodes
	label                float64                                                          // The label for this node - only used by Output Nodes
	myErr                float64                                                          // The error value for this node - only used by Output Nodes

	// Internal use

	dEdW        []float64 // The change in error by change in weight
	myoutCached bool      // Whether we have cached the output or not
	myout       float64   // My output
	backCached  bool      // Whether or not we have cached the back
	inputs      []float64 // The last input values into this node
	dEdOut      float64   //
	dOutdIn     float64   //
	dIndW       []float64 //
}

// newNode is a convenience constructor for nodes.
// This function returns a pointer to a new node whose values are set, with the exception of inputNodes
// and ouputNoddes, which are set to nil.
func newNode(wi weightInitializer, af activationFunction, da derivativeActivation, nodeIndex int, learningRate float64, layerIndex int, de derivativeError, nt nodeType) *node {
	return &node{
		inputNodes:           nil,
		outputNodes:          nil,
		weightInitializer:    wi,
		activation:           af,
		derivativeActivation: da,
		myIndex:              nodeIndex,
		learningRate:         learningRate,
		layerIndex:           layerIndex,
		derivativeError:      de,
		weights:              []float64{},
		myType:               nt,
		inputValue:           float64(0.0),
		label:                float64(0.0),
		myErr:                float64(0.0),
		dEdW:                 []float64{},
		myoutCached:          false,
		myout:                float64(0.0),
		backCached:           false,
		inputs:               []float64{},
		dEdOut:               float64(0.0),
		dOutdIn:              float64(0.0),
		dIndW:                []float64{},
	}
}

func (n *node) String() string {
	return fmt.Sprintf("NODE %d.%d", n.layerIndex, n.myIndex)
}

// initialize creates this node's input weights.
func (n *node) initialize() {
	for _, node := range n.inputNodes {
		weight := n.weightInitializer(n.layerIndex, node.myIndex, n.myIndex)
		n.weights = append(n.weights, weight)
	}
}

// updateWeights updates all the weights coming into this node.
func (n *node) updateWeights() {
	if n.myType == inputNode {
		// pass
	} else {
		for i := 0; i < len(n.weights); i++ {
			n.weights[i] -= n.learningRate * n.dEdW[i]
		}
	}
}

// invalidateCache invalidates anything that has been cached by this node.
func (n *node) invalidateCache() {
	n.myoutCached = false
	n.myout = float64(0.0)
	n.backCached = false
	n.inputs = []float64{}
}

// forward does a forward pass of whatever data is currently loaded into the network's input layer.
func (n *node) forward() float64 {
	if n.myoutCached {
		return n.myout
	}

	fmt.Fprintf(os.Stdout, "  Forward node %s\n", n.String())
	if n.myType == inputNode {
		n.myout = n.inputValue
	} else {
		n.myout = float64(0.0)
		for index, otherNode := range n.inputNodes {
			output := otherNode.forward()
			n.inputs = append(n.inputs, output)
			n.myout += n.weights[index] * output
		}
		n.myout += n.biasValue * n.biasWeight
		myoutActivated := n.activation(n.myout)
		n.myout = myoutActivated
	}

	n.myoutCached = true
	fmt.Fprintf(os.Stdout, "      Returning %f\n", n.myout)
	return n.myout
}

// loadError loads the error and label into the output node.
func (n *node) loadError(err float64, label float64) {
	n.label = label
	n.myErr = label
}

// backward calculates the gradients but does not update the weights.
// To update the weights, call n.UpdateWeights().
func (n *node) backward() (dEdOut float64, derAct float64, weights []float64) {
	fmt.Fprintf(os.Stdout, "Backproping node %s\n", n.String())
	if !n.backCached {
		if n.myType == outputNode {
			fmt.Fprintf(os.Stdout, "  Output Node\n")
			n.dEdOut = n.derivativeError(n.myout, n.label)
			fmt.Fprintf(os.Stdout, "  Calculated dEdOut as %f", n.dEdOut)
			if len(n.dEdW) == 0 {
				for _, input := range n.inputs {
					n.dEdW = append(n.dEdW, n.dEdOut*n.derivativeActivation(n.myout)*input)
				}
			} else {
				for i, input := range n.inputs {
					n.dEdW[i] = n.dEdOut * n.derivativeActivation(n.myout) * input
				}
			}
		} else {
			fmt.Fprintf(os.Stdout, "  Non-output node\n")
			n.dEdOut = float64(0.0)
			for _, l := range n.outputNodes {
				dEdInL, dInLdOut, ws := l.backward()
				n.dEdOut += dEdInL * dInLdOut * ws[n.myIndex]
			}
			n.dOutdIn = n.derivativeActivation(n.myout)
			n.dIndW = n.inputs
			if len(n.dEdW) == 0 {
				for _, dIndW := range n.dIndW {
					n.dEdW = append(n.dEdW, n.dEdOut*n.dOutdIn*dIndW)
				}
			} else {
				for index, dIndW := range n.dIndW {
					n.dEdW[index] = n.dEdOut * n.dOutdIn * dIndW
				}
			}
		}

		n.backCached = true
	}
	return n.dEdOut, n.dOutdIn, n.weights
}

// id returns a within-network unique string
func (n *node) id() string {
	var typestr string
	switch n.myType {
	case inputNode:
		typestr = "I"
	case hiddenNode:
		typestr = "H"
	case outputNode:
		typestr = "O"
	}
	return fmt.Sprintf("%s%d", typestr, n.myIndex)
}

// loadInput loads an input into this node
func (n *node) loadInput(invalue float64) {
	n.inputValue = invalue
}
