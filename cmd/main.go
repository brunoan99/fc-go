package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/brunoan99/fullcycle-go/internal/order/infra/database"
	usecase "github.com/brunoan99/fullcycle-go/internal/order/use-cases"
	"github.com/brunoan99/fullcycle-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	worker(out, uc, 1)
	// input := usecase.OrderInputDTO{
	// 	ID:    "1234",
	// 	Price: 100.0,
	// 	Tax:   10.0,
	// }
	// output, err := uc.Execute(input)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(output.FinalPrice)
}

func worker(develiveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range develiveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}
		input.Tax = 10.0
		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println("Error processing OrderInputDto", err)
		}
		fmt.Println("Recieved Message")
		msg.Ack(false)
	}
}
