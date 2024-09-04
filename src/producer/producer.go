package main

import (
	"context"
	"fmt"
	"golang-kafka-postgre/src/constant"
	"golang-kafka-postgre/src/interfaces"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaCon interfaces.KafkaWriter

func produce(ctx context.Context, msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recovered: %v\n", r)
		}
	}()

	start := time.Now()
	i := 1

	err := kafkaCon.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(i)),
		Value: []byte(msg),
	})

	if err != nil {
		panic("message failed to send " + err.Error())
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("writes message took %v seconds \n", elapsed)
}

func parseData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		message, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()
		go produce(ctx, string(message))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func initCon() {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{constant.BrokerAddress1, constant.BrokerAddress2, constant.BrokerAddress3},
		Topic:     constant.TopicCreateConnote,
		BatchSize: 15,
		Logger:    l,
	})

	kafkaCon = w
}

func main() {
	defer kafkaCon.Close()
	initCon()
	http.HandleFunc("/", parseData)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8153", nil); err != nil {
		log.Fatal(err)
	}
}
