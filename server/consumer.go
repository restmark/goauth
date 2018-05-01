package server

import (
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
)

func (a *API) SetupConsumer() {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// init consumer
	brokerList := a.Config.GetStringSlice("brokers")
	topics := []string{a.Config.GetString("kafka_topic")}
	consumer, err := cluster.NewConsumer(brokerList, a.Config.GetString("group_id"), topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				a.TopicRouter.CallHandler(msg.Topic, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			log.Println("Connection to the kafka server disconnected.")
		}
	}
}
