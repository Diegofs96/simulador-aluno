package main

import (
	"fmt"
	route2 "github.com/Diegofs96/simulador-aluno/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := gotdotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgchan := make(chan *ckafka.Message)
	consumer:= kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}