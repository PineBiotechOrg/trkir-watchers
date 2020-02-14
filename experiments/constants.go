package main

import (
    "os"
)

const (
    wsPort  = "8081"
    wsStreamingPath = "/streaming"
    separator = ","
    topic = "main_broker"
)

var rpiIps = os.Getenv("RPI_IPS")
var mouseIds = os.Getenv("MOUSE_IDS")
var kafkaServers = os.Getenv("KAFKA_SERVERS")
