package kafka

import (
	route2 "github.com/Diegofs96/simulador-aluno/application/route"
	ckafka "github.com/confluentinc/confluent-fakfa-go/kafka"
	"github.com/Diegofs96/simulador-aluno/infra/kafka"
	"log"
	"os"
)


func Produce(msg *ckafka.Message){
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}