package main

import "fmt"

// WeightInitializer is a typedef for functions passed around for initializing weights.
type WeightInitializer func(layerIndex int, myIndex int, otherNodeIndex int) (weight floatXX)

// ActivationFunction is a typedef for activation functions.
type ActivationFunction func(dotProduct floatXX) floatXX

// DerivativeActivation is a typedef of the derivative of an activation function.
type DerivativeActivation func(z floatXX) floatXX

// DerivativeError is a typedef of the derivative of an error function.
type DerivativeError func(thisNodeOutput floatXX, thisNodeLabel floatXX) floatXX

// Node is the main node object in a network.
type Node struct {
	inputNodes           []*Node                                                                // The list of nodes that input into this one
	outputNodes          []*Node                                                                // The list of nodes that this one input into
	weightInitializer    func(layerIndex int, myIndex int, otherNodeIndex int) (weight floatXX) // The function used to initialize the weights into this node
	activation           func(dotProduct floatXX) floatXX                                       // The activation function for this node
	derivativeActivation func(z floatXX) floatXX                                                // The derivative of the activation function
	myIndex              int                                                                    // The index of this node within a layer, 0-based
	learningRate         floatXX                                                                // The learning rate for this node's input weights
	biasValue            floatXX                                                                // The value of this layer's bias node
	biasWeight           floatXX                                                                // The initial weight of the bias node
	layerIndex           int                                                                    // The index of the layer that this node is a part of, 0-based
	derivativeError      func(thisNodeOutput floatXX, thisNodeLabel floatXX) floatXX            // The derivative of the error function
	weights              []floatXX                                                              // The weights, lined up with inputNodes
}

// NewNode is a convenience constructor for Nodes.
// This function returns a pointer to a new Node whose values are set, with the exception of inputNodes
// and ouputNoddes, which are set to nil.
func NewNode(weightInitializer WeightInitializer, activationFunction ActivationFunction, derivativeActivation DerivativeActivation, nodeIndex int, learningRate floatXX, layerIndex int, derivativeError DerivativeError) *Node {
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
		weights:              []floatXX{},
	}
}

func (n Node) String() string {
	return fmt.Sprintf("NODE %d.%d", n.myIndex, n.layerIndex)
}
