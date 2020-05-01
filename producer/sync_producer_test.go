package main

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
)

func messageChecker(result string) func([]byte) error {
	return func(val []byte) error {
		fmt.Println(string(val))
		fmt.Println(result)

		return nil
	}
}

// func generateRegexpChecker(re string) func([]byte) error {
// 	return func(val []byte) error {
// 		matched, err := regexp.MatchString(re, string(val))
// 		if err != nil {
// 			return errors.New("Error while trying to match the input message with the expected pattern: " + err.Error())
// 		}
// 		if !matched {
// 			return fmt.Errorf("No match between input value \"%s\" and expected pattern \"%s\"", val, re)
// 		}
// 		return nil
// 	}
// }

func TestSyncProducerShouldSendMessage(t *testing.T) {
	sp := mocks.NewSyncProducer(t, nil)
	defer func() {
		if err := sp.Close(); err != nil {
			t.Error(err)
		}
	}()

	// sp.ExpectSendMessageAndSucceed()
	sp.ExpectSendMessageWithCheckerFunctionAndSucceed(messageChecker("test"))

	msg := &sarama.ProducerMessage{
		Topic: "test topic",
		Value: sarama.StringEncoder("test value"),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("traceparent"),
				Value: []byte("traceparent value"),
			},
		},
	}
	fmt.Println(msg.Topic)
	fmt.Println(msg.Value)
	fmt.Println(string(msg.Headers[0].Key))
	fmt.Println(string(msg.Headers[0].Value))
	sp.SendMessage(msg)
}
