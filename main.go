package main

import (
	"fmt"
	route2 "github.com/Diegofs96/simulador-aluno/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err :=gotdotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}