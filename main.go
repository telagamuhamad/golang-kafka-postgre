package main

import (
	"fmt"
	"golang-kafka-postgre/src/database"
	"golang-kafka-postgre/src/kafka"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	go kafka.StartConsumer()

	fmt.Println("Consumer started")
	select {} // block forever
}
