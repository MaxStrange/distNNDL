package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func startLabeledDataTopic(broker string, sigChan <-chan os.Signal, mainChan chan<- string) {
	dataConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        "labeledDataGroup",
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(-1)
	}

	err = dataConsumer.SubscribeTopics([]string{"labeledDataTopic"}, nil)
	run := true
	for run == true {
		select {
		case _ = <-sigChan:
			fmt.Println("Terminating.")
			run = false
		case ev := <-dataConsumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				dataConsumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				dataConsumer.Unassign()
			case *kafka.Message:
				mainChan <- string(e.Value)
			case kafka.PartitionEOF:
				fmt.Fprintf(os.Stderr, "%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				mainChan <- "TERM"
				run = false
			}
		}
	}
}
