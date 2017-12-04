package main

import "fmt"

type NodeType int

const (
	InputNode NodeType = iota
	HiddenNode
	OutputNode
)

// WeightInitializer is a typedef for functions passed around for initializing weights.
type WeightInitializer func(layerIndex int, inIndex int, outIndex int) (weight FloatXX)

// ActivationFunction is a typedef for activation functions.
type ActivationFunction func(dotProduct FloatXX) FloatXX

// DerivativeActivation is a typedef of the derivative of an activation function.
type DerivativeActivation func(z FloatXX) FloatXX

// DerivativeError is a typedef of the derivative of an error function.
type DerivativeError func(thisNodeOutput FloatXX, thisNodeLabel FloatXX) FloatXX

// Node is the main node object in a network.
type Node struct {
	inputNodes           []*Node                                                          // The list of nodes that input into this one
	outputNodes          []*Node                                                          // The list of nodes that this one input into
	weightInitializer    func(layerIndex int, inIndex int, outIndex int) (weight FloatXX) // The function used to initialize the weights into this node
	activation           func(dotProduct FloatXX) FloatXX                                 // The activation function for this node
	derivativeActivation func(z FloatXX) FloatXX                                          // The derivative of the activation function
	myIndex              int                                                              // The index of this node within a layer, 0-based
	learningRate         FloatXX                                                          // The learning rate for this node's input weights
	biasValue            FloatXX                                                          // The value of this layer's bias node
	biasWeight           FloatXX                                                          // The initial weight of the bias node
	layerIndex           int                                                              // The index of the layer that this node is a part of, 0-based
	derivativeError      func(thisNodeOutput FloatXX, thisNodeLabel FloatXX) FloatXX      // The derivative of the error function
	weights              []FloatXX                                                        // The weights, lined up with inputNodes
	myType               NodeType                                                         // The type of this node
	inputValue           FloatXX                                                          // The input value into this node - only used by Input Nodes
	label                FloatXX                                                          // The label for this node - only used by Output Nodes
	myErr                FloatXX                                                          // The error value for this node - only used by Output Nodes

	// Internal use

	dEdW        []FloatXX // The change in error by change in weight // internal use
	myoutCached bool      // Whether we have cached the output or not
	myout       FloatXX   // My output
	backCached  bool      // Whether or not we have cached the back
	inputs      []FloatXX // The last input values into this node
	dEdOut      FloatXX   //
	dOutdIn     FloatXX   //
	dIndW       []FloatXX //
}

// NewNode is a convenience constructor for Nodes.
// This function returns a pointer to a new Node whose values are set, with the exception of inputNodes
// and ouputNoddes, which are set to nil.
func NewNode(weightInitializer WeightInitializer, activationFunction ActivationFunction, derivativeActivation DerivativeActivation, nodeIndex int, learningRate FloatXX, layerIndex int, derivativeError DerivativeError, nodeType NodeType) *Node {
	return &Node{
		inputNodes:           nil,
		outputNodes:          nil,
		weightInitializer:    weightInitializer,
		activation:           activationFunction,
		derivativeActivation: derivativeActivation,
		myIndex:              nodeIndex,
		learningRate:         learningRate,
		layerIndex:           layerIndex,
		derivativeError:      derivativeError,
		weights:              []FloatXX{},
		myType:               nodeType,
		inputValue:           FloatXX(0.0),
		label:                FloatXX(0.0),
		myErr:                FloatXX(0.0),
		dEdW:                 []FloatXX{},
		myoutCached:          false,
		myout:                FloatXX(0.0),
		backCached:           false,
		inputs:               []FloatXX{},
		dEdOut:               FloatXX(0.0),
		dOutdIn:              FloatXX(0.0),
		dIndW:                []FloatXX{},
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("NODE %d.%d", n.myIndex, n.layerIndex)
}

func (n *Node) Initialize() {
	for _, node := range n.inputNodes {
		weight := n.weightInitializer(n.layerIndex, node.myIndex, n.myIndex)
		n.weights = append(n.weights, weight)
	}
}

// UpdateWeights updates all the weights coming into this node.
func (n *Node) UpdateWeights() {
	if n.myType == InputNode {
		// pass
	} else {
		for i := 0; i < len(n.weights); i++ {
			n.weights[i] -= n.learningRate * n.dEdW[i]
		}
	}
}

// InvalidateCache invalidates anything that has been cached by this node.
func (n *Node) InvalidateCache() {
	n.myoutCached = false
	n.myout = FloatXX(0.0)
	n.backCached = false
	n.inputs = []FloatXX{}
}

// Forward does a forward pass of whatever data is currently loaded into the network's input layer.
func (n *Node) Forward() FloatXX {
	if n.myoutCached {
		return n.myout
	}

	if n.myType == InputNode {
		n.myout = n.inputValue
	} else {
		n.myout = FloatXX(0.0)
		for index, otherNode := range n.inputNodes {
			output := otherNode.Forward()
			n.inputs = append(n.inputs, output)
			n.myout += n.weights[index] * output
		}
		n.myout += n.biasValue * n.biasWeight
		myoutActivated := n.activation(n.myout)
		n.myout = myoutActivated
	}

	n.myoutCached = true
	return n.myout
}

// LoadError loads the error and label into the output node.
func (n *Node) LoadError(err FloatXX, label FloatXX) {
	n.label = label
	n.myErr = label
}

// Backward calculates the gradients but does not update the weights.
// To update the weights, call n.UpdateWeights().
func (n *Node) Backward() (dEdOut FloatXX, derAct FloatXX, weights []FloatXX) {
	if !n.backCached {
		if n.myType == OutputNode {
			n.dEdOut = n.derivativeError(n.myout, n.label)
			for i, input := range n.inputs {
				n.dEdW[i] = n.dEdOut * n.derivativeActivation(n.myout) * input
			}
		} else {
			n.dEdOut = FloatXX(0.0)
			for _, l := range n.outputNodes {
				dEdInL, dInLdOut, ws := l.Backward()
				n.dEdOut += dEdInL * dInLdOut * ws[n.myIndex]
			}
			n.dOutdIn = n.derivativeActivation(n.myout)
			n.dIndW = n.inputs
			for index, dIndW := range n.dIndW {
				n.dEdW[index] = n.dEdOut * n.dOutdIn * dIndW
			}
		}

		n.backCached = true
	}
	return n.dEdOut, n.dOutdIn, n.weights
}

// Id returns a within-network unique string
func (n *Node) Id() string {
	var typestr string
	switch n.myType {
	case InputNode:
		typestr = "I"
	case HiddenNode:
		typestr = "H"
	case OutputNode:
		typestr = "O"
	}
	return fmt.Sprintf("%s%d", typestr, n.myIndex)
}

// LoadInput loads an input into this node
func (n *Node) LoadInput(invalue FloatXX) {
	n.inputValue = invalue
}
