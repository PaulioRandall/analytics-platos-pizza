package database

import (
	"fmt"
	"time"
)

// Order represents an order of pizzas, one or many pizzas per order
type Order struct {
	// Unique identifier for each order placed by a table
	Id int

	// Date & time the order was placed
	// (entered into the system prior to cooking & serving)
	Datetime time.Time
}

func PrintOrders(orders []Order) {
	fmt.Println("[Orders]")
	fmt.Println(`"Id", "Datetime"`)
	for _, entry := range orders {
		fmt.Printf("%q, %q\n", entry.Id, entry.Datetime)
	}
}
