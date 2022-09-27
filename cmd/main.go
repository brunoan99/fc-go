package main

import (
	"database/sql"
	"fmt"

	"github.com/brunoan99/fullcycle-go/internal/order/infra/database"
	usecase "github.com/brunoan99/fullcycle-go/internal/order/use-cases"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100.0,
		Tax:   10.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output.FinalPrice)
}
