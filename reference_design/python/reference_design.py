"""
This is a quick-n-dirty reference design written up and debugged in single-threaded python
to make sure all the math works.
"""
import math
import random
import matplotlib.pyplot as plt

class Layer:
    def __init__(self, num, weight_initializer, activation, derivative_activation, learning_rate, layer_index, bias, bias_weight, derivative_error, input_layer=False, output_layer=False):
        """
        :param num:                     The number of nodes in this layer.
        :param weight_initializer:      fn(layer_index, node_from, node_to) -> initial weight
        :param activation:              fn(float) -> float (must be nonlinear and differentiable)
        :param derivative_activation:   fn(float) -> float (must be the derivative of the activation function)
        :param learning_rate:           The learning rate - a constant float, usually between 0.1 and 0.4
        :param layer_index:             The index of this layer in the network - 0 based.
        :param bias:                    The bias value for this layer.
        :param bias_weight:             The weight for the bias in this layer.
        :param derivative_error:        The derivative of the error function used on the output layer.
        :param input_layer:             True if this is the input layer.
        :param output_layer:            True if this is the output layer.
        """
        if input_layer:
            self.nodes = [InputNode(None, weight_initializer, activation, None, derivative_activation, i, learning_rate, None, None, layer_index, derivative_error) for i in range(num)]
        elif output_layer:
            self.nodes = [OutputNode(None, weight_initializer, activation, None, derivative_activation, i, learning_rate, bias, bias_weight, layer_index, derivative_error) for i in range(num)]
        else:
            self.nodes = [HiddenNode(None, weight_initializer, activation, None, derivative_activation, i, learning_rate, bias, bias_weight, layer_index, derivative_error) for i in range(num)]

    def set_previous_layer(self, layer):
        """
        Sets the layer that this layer is connected FROM

        :param layer: A Layer instance.
        """
        for n in self.nodes:
            n.input_nodes = layer.nodes

    def set_next_layer(self, layer):
        """
        Sets the layer that this layer connects TO

        :param layer: A Layer instance.
        """
        for n in self.nodes:
            n.output_nodes = layer.nodes

class Network:
    def __init__(self, input_layer, hidden_layers, output_layer, err_func):
        """
        :param input_layer: The input layer
        :param hidden_layers: An iterable of Layer objects to form the hidden layers, in order
        :param output_layer: The output layer
        :param err_func: fn([labels], [outputs]) -> [errors]
        """
        self.input_layer = input_layer
        self.hidden_layers = hidden_layers
        self.output_layer = output_layer
        self.err_func = err_func
        for n in self.all_neurons():
            n.initialize()

    def run(self, input_vectors, labels, epochs=1, batch_size=1):
        """
        Trains the network.

        :param input_vectors: The input dataset
        :param labels:        The label dataset, in the same order as input_vectors
        :param epochs:        The number of times to train the network over the whole dataset
        :param batch_size:    The number of items from input_vectors to draw and run before we train the network
        """
        assert len(input_vectors) == len(labels)
        assert batch_size <= len(input_vectors)
        assert len(input_vectors) > 0
        assert len(labels[0]) == len(self.output_layer.nodes)
        assert len(input_vectors[0]) == len(self.input_layer.nodes)

        accuracies = []
        for epoch_index in range(epochs):
            print("----------------------------------")
            print("EPOCH", epoch_index, ":")
            print("----------------------------------")
            for vec, label in zip(input_vectors, labels):
                print("DATA:", vec, "LABEL", label)
                for n, input_val in zip(self.input_layer.nodes, vec):
                    n.load_input(input_val)
                y_hat = [n.forward() for n in self.output_layer.nodes]
                err = self.err_func(label, y_hat)
                for n, e in zip(self.output_layer.nodes, err):
                    n.load_err(e, label[n.my_index])
                for n in self.input_layer.nodes:
                    n.backward()
                for n in self.all_neurons():
                    n.update_weights()
                    n.invalidate_cache()
            if epoch_index % 10 == 0:
                accuracies.append(self.evaluate(input_vectors, labels, log=False))
        plt.plot(accuracies)
        plt.title("Accuracy over Time")
        plt.show()

    def evaluate(self, input_vectors, labels, log=True):
        """
        Evaluate the network.

        :param input_vectors: The vectors to evaluate on
        :param labels:        The label vectors in the same order as input_vectors
        :param log:           If True, print stuff out
        """
        if log:
            print("!!!!!!!!!!!!!!!!!! ACCURACY !!!!!!!!!!!!!!!!!!!!!")
            print("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
        total_acc = 0.0
        for vec, label in zip(input_vectors, labels):
            for n, input_val in zip(self.input_layer.nodes, vec):
                n.load_input(input_val)
            y_hat = [n.forward() for n in self.output_layer.nodes]
            acc = int(round(y_hat[0])) == int(round(label[0]))
            total_acc += acc
            for n in self.all_neurons():
                n.invalidate_cache()
            if log:
                print(":::::::::::::::::::::::")
                print("Input:", vec)
                print("Label:", label)
                print("Output:", int(round(y_hat[0])))
                print(":::::::::::::::::::::::")
        total_acc /= len(input_vectors)
        if log:
            print("!!!!!!!!!!!!!!!!!! ACCURACY !!!!!!!!!!!!!!!!!!!!!")
            print("ACCURACY:", str(total_acc * 100) + "%")
            print("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
        return total_acc

    def all_neurons(self):
        """
        Iterator to yield each neuron in order.
        """
        for n in self.input_layer.nodes:
            yield n
        for h in self.hidden_layers:
            for n in h.nodes:
                yield n
        for n in self.output_layer.nodes:
            yield n

class Node(object):
    def __init__(self, input_nodes, weight_initializer, activation, output_nodes, derivative_activation, my_index, learning_rate, bias, bias_weight, layer, derivative_error):
        """
        Abstract base class

        :param input_nodes:             The list of nodes that input into this one
        :param weight_initializer:      The function used to initialize the weights
        :param activation:              The activation function
        :param output_nodes:            The list of nodes that this one inputs into
        :param derivative_activation:   The derivative of the activation function
        :param my_index:                This node's unique index within a layer, 0-based
        :param learning_rate:           The learning rate for this node
        :param bias:                    The bias input to this node
        :param bias_weight:             The weight of the bias node
        :param layer:                   The layer index that we belong to
        :param derivative_error:        The derivative of the error function
        """
        self.input_nodes = input_nodes
        self._inputs = []
        self.activation = activation
        self.output_nodes = output_nodes
        self._myout = None
        self._myout_cached = False
        self.derivative_activation = derivative_activation
        self._back_cached = False
        self.learning_rate = learning_rate
        self.weight_initializer = weight_initializer
        self.my_index = my_index
        self.bias = bias
        self.bias_weight = bias_weight
        self.layer = layer
        self.derivative_error = derivative_error

    def initialize(self):
        """
        Initialize the Node

        self.weights is initialized. This field is the weight coming from each input into this node.
        """
        if self.input_nodes is not None:
            self.weights = [self.weight_initializer(self.layer, i.my_index, self.my_index) for i in self.input_nodes]
        else:
            self.weights = []

    def update_weights(self):
        """
        Updates the weights going INTO this node.
        """
        for index in range(len(self.weights)):
            self.weights[index] -= self.learning_rate * self._dE_dW[index]

    def invalidate_cache(self):
        """
        Invalidates this node's cache.
        """
        self._myout_cached = False
        self._myout = None
        self._back_cached = False
        self._inputs = []

    def forward(self):
        """
        Calculate the forward pass.
        """
        if self._myout_cached:
            return self._myout
        else:
            myout = 0.0
            for index, i in enumerate(self.input_nodes):
                output = i.forward()
                self._inputs.append(output)
                myout += self.weights[index] * output
            myout += self.bias * self.bias_weight
            myout_activated = self.activation(myout)
            self._myout = myout_activated
            self._myout_cached = True
            return myout_activated

    def backward(self):
        """
        Calculate all the partial derivatives so that you can call update_weights().
        """
        if not self._back_cached:
            self._dE_dOut = 0.0
            for index, l in enumerate(self.output_nodes):
                dE_dIn_l, dIn_l_dOut, ws = l.backward()
                self._dE_dOut += dE_dIn_l * dIn_l_dOut * ws[self.my_index]
            self._dOut_dIn = self.derivative_activation(self._myout)
            self._dIn_dW = self._inputs
            self._dE_dW = [self._dE_dOut * self._dOut_dIn * dIndW for dIndW in self._dIn_dW]
            self._back_cached = True
        return self._dE_dOut, self.derivative_activation(self._myout), self.weights

    def id(self):
        return "NONE"


class InputNode(Node):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self._input_val = None

    def load_input(self, inval):
        """
        Load the input value for this forward pass.
        """
        self._input_val = inval

    def forward(self):
        """
        Do the forward pass - just return my input value.
        """
        self._myout = self._input_val
        self._myout_cached = True
        return self._myout

    def id(self):
        return "I" + str(self.my_index)

    def update_weights(self):
        """
        Input nodes do not have weights.
        """
        pass

class OutputNode(Node):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def load_err(self, err, label):
        """
        Load the error value for this Node.

        :param err: The error value
        :param label: The label for this node
        """
        self.label = label

    def backward(self):
        """
        """
        if not self._back_cached:
            self._dE_dOut = self.derivative_error(self._myout, self.label)
            self._dE_dW = [self._dE_dOut * self.derivative_activation(self._myout) * inp for inp in self._inputs]
            self._back_cached = True
        return self._dE_dOut, self.derivative_activation(self._myout), self.weights

    def id(self):
        return "O" + str(self.my_index)

class HiddenNode(Node):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

    def id(self):
        return "H" + str(self.my_index)






def small_random_numbers(_layer, _out_index, _in_index):
    """
    Weight initializer for small random values
    """
    return random.uniform(-0.3, 0.3)

def blog_init(layer, out_index, in_index):
    """
    Initializer borrowed to check work against a blog post.
    """
    if layer == 0:
        assert False
    elif layer == 1:
        if out_index == 0 and in_index == 0:
            return 0.15
        elif out_index == 0 and in_index == 1:
            return 0.25
        elif out_index == 1 and in_index == 0:
            return 0.2
        elif out_index == 1 and in_index == 1:
            return 0.3
        else:
            assert False
    elif layer == 2:
        if out_index == 0 and in_index == 0:
            return 0.4
        elif out_index == 0 and in_index == 1:
            return 0.45
        elif out_index == 1 and in_index == 0:
            return 0.5
        elif out_index == 1 and in_index == 1:
            return 0.55
        else:
            assert False


def sigmoid(z):
    """
    Logistic activation function
    """
    activation = 1.0 / (1.0 + math.e ** -z)
    return activation

def dersigmoid(z):
    """
    Derivitive of the logistic function
    """
    der = z * (1 - z)
    return der

def error_function(labels, y_hat):
    """
    Error function
    """
    err = [0.5 * (t - y) ** 2 for t, y in zip(labels, y_hat)]
    return err

def dererror(y_hat, label):
    """
    The derivitive of the error function for a single node (not a vector).
    """
    return y_hat - label


INPUT_VECTORS_AND_SET = [ [0, 0], [1, 0], [0, 1], [1, 1] ]
LABEL_VECTORS_AND_SET = [ [0],    [0],    [0],    [1]    ]

INPUT_VECTORS_BLOG_SET = [ [0.05, 0.10] ]
LABEL_VECTORS_BLOG_SET = [ [0.01, 0.99] ]

INPUT_VECTORS_XOR_SET = [ [0, 0], [1, 0], [0, 1], [1, 1] ]
LABEL_VECTORS_XOR_SET = [ [0],    [1],    [1],    [0]    ]

if __name__ == "__main__":
    random.seed(541908)

    # AND
    input_vectors  = INPUT_VECTORS_XOR_SET
    labels = LABEL_VECTORS_XOR_SET

    output_layer = Layer(num=1,
                        weight_initializer=small_random_numbers,
                        activation=sigmoid,
                        derivative_activation=dersigmoid,
                        learning_rate=0.6,
                        layer_index = 2,
                        bias=1.0,
                        bias_weight=0.6,
                        derivative_error=dererror,
                        output_layer=True
                      )

    hidden_layer = Layer(num=2,
                        weight_initializer=small_random_numbers,
                        activation=sigmoid,
                        derivative_activation=dersigmoid,
                        learning_rate=0.6,
                        layer_index = 1,
                        bias=1.0,
                        bias_weight=0.35,
                        derivative_error=dererror
                      )

    input_layer = Layer(num=2,
                        weight_initializer=small_random_numbers,
                        activation=sigmoid,
                        derivative_activation=dersigmoid,
                        learning_rate=0.6,
                        layer_index = 0,
                        bias=None,
                        bias_weight=None,
                        derivative_error=dererror,
                        input_layer=True
                      )

    output_layer.set_previous_layer(hidden_layer)
    input_layer.set_next_layer(hidden_layer)
    hidden_layer.set_previous_layer(input_layer)
    hidden_layer.set_next_layer(output_layer)
    network = Network(input_layer, [hidden_layer], output_layer, error_function)

    network.run(input_vectors, labels, epochs=20000000)
    network.evaluate(input_vectors, labels)

