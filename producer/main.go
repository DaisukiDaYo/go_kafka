package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"

	"time"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "Traceparent"
)

var TraceID [16]byte = [16]byte{'0', '2', '0', '0', '0', '0', '4'}

type response1 struct {
	Page   int
	Fruits []string
}

func main() {
	fmt.Printf("%T\n", TraceID[:])
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	// read command line input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish without goroutine
		publish(msg, producer)

		// publish with go routine
		// go publish(msg, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	version, err := sarama.ParseKafkaVersion("2.0.1")
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	// producer config
	config := sarama.NewConfig()
	config.Version = version
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("trace_id"),
				Value: []byte("0000-75c4e49c781d8f63db39535fe203c945-cce5ec01ecf0f2c1-0000"),
			},
		},
	}

	now := time.Now().UTC()
	filename := fmt.Sprintf("dev_reverse_%s_%d%s%d_podid.txt", strings.ToLower(now.Format("Mon")), now.Year(), now.Format("01"), now.Day())
	f, err := os.OpenFile("../kafka-error-logs/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1)

	w := bufio.NewWriter(f)

	//msg.Headers["1", "2"]

	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	file, err := w.WriteString(string(res1B) + "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println("file", file)
	w.Flush()

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
