package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type producer struct {
	kafkaProducer *kafka.Producer
}

func makeProducer(broker string) *producer {
	props := &kafka.ConfigMap{"bootstrap.servers": broker}
	p, err := kafka.NewProducer(props)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start producer: %s\n", err)
		os.Exit(-1)
	}
	return &producer{
		kafkaProducer: p,
	}
}

func (p *producer) startProducer(sigChan <-chan os.Signal) {
	run := true
	for run {
		select {
		case _ = <-sigChan:
			fmt.Println("Terminating.")
			run = false
		case e := <-p.kafkaProducer.Events():
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Fprintf(os.Stderr, "Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message %v to topic %s [%d] at offset %v\n", m, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}
	p.kafkaProducer.Close()
}

func (p *producer) send(key string, msg string, topic string) {
	m := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(msg),
		Key:   []byte(key),
	}
	fmt.Printf("Sending message to %s: K: %s V: %s\n", topic, key, msg)
	fmt.Printf("Message as internal representation: %v", m)
	p.kafkaProducer.ProduceChannel() <- m
}
