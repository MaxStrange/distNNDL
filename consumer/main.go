// This is the main file for a back-end node
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("USAGE: %s <UNIQUE_GROUP_ID> <BROKER_IP:BROKER_PORT>", os.Args[0])
		os.Exit(1)
	}

	group := os.Args[1]
	broker := os.Args[2]

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	metaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"}})

	if err != nil {
		fmt.Println("Failed to create consumer: %s", err)
		os.Exit(-1)
	}

	// TODO: run this next part in its own thread and also spawn another thread that subscribes to labeledDataTopic
	err = metaConsumer.SubscribeTopics([]string{"NNDLTopic"}, nil) //TOOD: look up the args for this
	run := true
	for run == true {
		select {
		case _ = <-sigchan:
			fmt.Println("Terminating.")
			run = false
		case ev := <-metaConsumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Println("Assigned partition: %% %v", e)
				metaConsumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Println("%% %v", e)
				metaConsumer.Unassign()
			case *kafka.Message:
				fmt.Println("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
			case kafka.PartitionEOF:
				fmt.Println("%% Reached %v", e)
			case kafka.Error:
				fmt.Println("%% Error: %v", e)
				run = false
			}
		}
	}
}
