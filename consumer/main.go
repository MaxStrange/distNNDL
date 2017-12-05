// This is the main file for a back-end node
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "USAGE: %s <UNIQUE_GROUP_ID> <BROKER_IP:BROKER_PORT>\n", os.Args[0])
		os.Exit(1)
	}

	group := os.Args[1]
	broker := os.Args[2]

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	nndlChan := make(chan string, 1)
	dataChan := make(chan string, 5)
	trainingChan := make(chan string, 1)

	go startnndlConsumer(group, broker, sigChan, nndlChan)
	go startLabeledDataTopic(broker, sigChan, dataChan)
	kafkaProducer := makeProducer(broker)
	go kafkaProducer.startProducer(sigChan)

	batchLength, epochLength, mixLength, jobID := getMetaInformation(broker, group, sigChan)
	net := makeNetwork(batchLength, epochLength, mixLength)
	run := true
	for run == true {
		select {
		case <-sigChan:
			fmt.Println("Terminating.")
			run = false
		case msg := <-nndlChan:
			fmt.Println(msg)
			net.parseNetFromNNDL(msg)
			go net.train(dataChan, trainingChan)
		case msg := <-trainingChan:
			fmt.Println("Sending:\n" + msg)
			kafkaProducer.send(jobID, msg, "statusInfo")
		}
	}
}
