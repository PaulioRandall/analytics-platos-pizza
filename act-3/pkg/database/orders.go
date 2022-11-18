package database

import (
	"fmt"
	"time"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
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
		fmt.Printf("%d, %q\n", entry.Id, entry.Datetime)
	}
}

func QueryPrintOrders(db PlatosPizzaDatabase) error {
	records, e := db.QueryHeadOrders()

	if e != nil {
		return err.Wrap(e, "Quering all metadata")
	}

	PrintOrders(records)
	fmt.Println("...")

	return nil
}
