package database

import (
	"fmt"

	"github.com/PaulioRandall/trackable-go"
)

type Pizza struct {
	// Unique identifier for each pizza (constituted by its type and size)
	Id string

	// Foreign key that ties each pizza to its broader pizza type
	PizzaTypeId string

	// Size of the pizza (Small, Medium, Large, X Large, or XX Large)
	Size string

	// Price of the pizza in USD
	Price float64
}

func PrintPizzas(pizzas []Pizza) {
	fmt.Println("[Pizzas]")
	fmt.Println(`"ID", "Pizza Type ID", "Size", "Price"`)
	for _, v := range pizzas {
		fmt.Printf("%q, %q, %q, %.2f\n", v.Id, v.PizzaTypeId, v.Size, v.Price)
	}
}

func QueryPrintPizzas(db PlatosPizzaDatabase) error {
	records, e := db.HeadPizzas()

	if e != nil {
		return trackable.Wrap(e, "Quering head of pizzas")
	}

	PrintPizzas(records)
	fmt.Println("...")

	return nil
}
