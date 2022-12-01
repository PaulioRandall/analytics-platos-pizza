package database

import (
	"fmt"

	"github.com/PaulioRandall/trackable"
)

type Pizza struct {
	// Unique identifier for each pizza (constituted by its type and size)
	Id string

	// Foreign key that ties each pizza to its broader pizza type
	TypeId string

	// Size of the pizza (Small, Medium, Large, X Large, or XX Large)
	Size string

	// Price of the pizza in USD
	Price float64
}

func PrintPizzas(pizzas []Pizza) {
	fmt.Println("[Pizzas]")
	fmt.Println(`"ID", "Type ID", "Size", "Price"`)
	for _, v := range pizzas {
		fmt.Printf("%q, %q, %q, %.2f\n", v.Id, v.TypeId, v.Size, v.Price)
	}
}

func QueryPrintPizzas(db PlatosPizzaDatabase) error {
	records, e := db.HeadPizzas()

	if e != nil {
		return trackable.WrapAtInterface(e, "database.QueryPrintPizzas")
	}

	PrintPizzas(records)
	fmt.Println("...")

	return nil
}
