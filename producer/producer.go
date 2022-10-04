package producer

import (
	"encoding/hex"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers = []string{"localhost:9093"}
	KafkaTopic   = "sarama_topic"
	enqueued     int
)

// StartProducer runs the AsyncProducer
func StartProducer() (sarama.AsyncProducer, chan os.Signal){

	producer, err := setupProducer()
	if err != nil {
		panic(err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// produceMessages(producer, signals)

	// log.Printf("Kafka AsyncProducer finished with %d messages produced.", enqueued)
	return producer, signals
}

// setupProducer will create a AsyncProducer and returns it
func setupProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	return sarama.NewAsyncProducer(kafkaBrokers, config)
}

// produceMessages will send 'testing 123' to KafkaTopic each second, until receive a os signal to stop e.g. control + c
// by the user in terminal
func ProduceMessages(producer sarama.AsyncProducer, signals chan os.Signal, topic string, msg string) {
	// for {
	// 	time.Sleep(time.Second)
		// valueBytes := []byte(time.Now().Format("15:04:05.000"))
	// 	valueHash := sha256.Sum256(valueBytes)
	// 	valueString := hex.EncodeToString(valueHash[:])
	// 	message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder(valueString)}
	// 	select {
	// 	case producer.Input() <- message:
	// 		enqueued++
	// 		log.Println("New Message produced")
	// 	// case <-signals:
	// 	// 	producer.AsyncClose() // Trigger a shutdown of the producer.
	// 	// 	return
	// 	}
	// }
	valueBytes := []byte(msg)
	valueString := hex.EncodeToString(valueBytes)
	message := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(valueString), Key: sarama.StringEncoder("url")}
	select {
	case producer.Input() <- message:
		enqueued++
		log.Println("New Message produced")
	// case <-signals:
	// 	producer.AsyncClose() // Trigger a shutdown of the producer.
	// 	return
	}
}
