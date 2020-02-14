package main

import (
    "fmt"

    "github.com/gorilla/websocket"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func readRpi(rpiIp string, mouseId string) {
    producerTopic := topic
    // create producer
    producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServers})
    if err != nil {
        panic(err)
    }
    defer producer.Close()


    rpiUrl := fmt.Sprintf("ws://%s:%s%s", rpiIp, wsPort, wsStreamingPath)

    fmt.Printf("%s\n", rpiUrl)

    c, _, err := websocket.DefaultDialer.Dial(rpiUrl, nil)
    if err != nil {
        fmt.Println("Can't connect to ws server : ", err)
        return
    }

    for {
        _, message, err := c.ReadMessage()

        if err != nil {
            fmt.Println("read err:", err)
            return
        }

        producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &producerTopic, Partition: kafka.PartitionAny},
            Value: message,
            Key: []byte(mouseId),
        }, nil)
    }
}
