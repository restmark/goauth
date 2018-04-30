package server

import (
	"github.com/Shopify/sarama"
	"github.com/restmark/goauth/services"
)

func (a *API) SetupProducer() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	//tlsConfig := createTlsConfiguration()
	//if tlsConfig != nil {
	//	config.Net.TLS.Config = tlsConfig
	//	config.Net.TLS.Enable = true
	//}

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.
	brokerList := a.Config.GetStringSlice("brokers")
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return err
	}

	kafkaService := services.NewKafka(producer)

	a.Kafka = kafkaService

	return nil
}
