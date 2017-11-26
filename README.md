# Distributed Neural Network Training using Neural Network Description Language

This is a final project for a distributed computing class I am taking as part
of my master's in computer science at the University of Washington Bothell.

It is a multi-tier server architecture that takes in three files: a .ndl file, and two .csv files:
one for input vectors and one for labels. It then trains the network described in the .ndl file
in a distributed way, using Kafka and Go.

## NNDL

The .ndl file must match the following specifications:

*TODO*

## CSV

The two .csv files work as follows:

### Input Vectors

The .csv file that contains the training data must be of the form:

```csv
input_node_0's value, input_node_1's value, etc.
input_node_0's value, input_node_1's value, etc.
input_node_0's value, input_node_1's value, etc.
etc.
```

Where each line in the .csv file is a feature vector to be input into the network for a forward pass, and each
value within a line is a value in the feature vector.

### Labels

The .csv file that contains the labels for the training data must be of the form:

```csv
output_node_0's correct value, output_node_1's correct value, etc.
output_node_0's correct value, output_node_1's correct value, etc.
output_node_0's correct value, output_node_1's correct value, etc.
etc.
```

Where each line in the .csv file is an output vector - a vector formed via the expected/correct outputs for
each output node, and each value within a line is a value in the output vector.

There must be exactly the same number of lines in each .csv file - one label for each input.
