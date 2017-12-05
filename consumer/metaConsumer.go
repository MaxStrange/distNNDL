package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func getMetaInformation(broker string, groupID string, sigChan <-chan os.Signal) (batchLength int, epochLength int, mixLength int, jobID string) {
	metaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        groupID,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(-1)
	}

	err = metaConsumer.SubscribeTopics([]string{"metaDataTopic"}, nil)
	var batchLengthStr, epochLengthStr, mixLengthStr string
	run := true
	for run == true {
		select {
		case _ = <-sigChan:
			fmt.Println("Terminating.")
			run = false
		case ev := <-metaConsumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				metaConsumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				metaConsumer.Unassign()
			case *kafka.Message:
				switch string(e.Key) {
				case "batchSize":
					fmt.Fprintf(os.Stdout, "BatchSize: %s\n", string(e.Value))
					batchLengthStr = string(e.Value)
				case "epochLength":
					fmt.Fprintf(os.Stdout, "EpochLength: %s\n", string(e.Value))
					epochLengthStr = string(e.Value)
				case "mixLength":
					fmt.Fprintf(os.Stdout, "MixLength: %s\n", string(e.Value))
					mixLengthStr = string(e.Value)
				case "jobID":
					fmt.Fprintf(os.Stdout, "JobID: %s\n", string(e.Value))
					jobID = string(e.Value)
				case "terminate":
					fmt.Fprintf(os.Stdout, "Received TERM for metadata\n")
					run = false
				default:
					fmt.Fprintf(os.Stderr, "Got unexpected metainformation: key: %s", string(e.Key))
					run = false
				}
			case kafka.PartitionEOF:
				fmt.Fprintf(os.Stderr, "%% Reached %v\n", e)
				if mixLengthStr != "" && epochLengthStr != "" && batchLengthStr != "" {
					run = false
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "Kafka gave an error %v\n", e)
				os.Exit(-1)
			}
		}
	}

	ints := []int{}
	strings := []string{batchLengthStr, epochLengthStr, mixLengthStr}
	for _, s := range strings {
		fmt.Println(s)
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Got non-integer value %s when parsing metadata.\n", s)
			os.Exit(-1)
		}
		ints = append(ints, int(i))
	}
	fmt.Fprintf(os.Stdout, "Integers: %d %d %d\n", ints[0], ints[1], ints[2])
	return ints[0], ints[1], ints[2], jobID
}
