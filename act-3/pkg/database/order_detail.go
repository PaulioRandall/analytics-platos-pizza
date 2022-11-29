package database

import (
	"fmt"

	"github.com/PaulioRandall/trackable-go"
)

// OrderDetail represents a specific pizza type order, one or more pizzas,
// within an order
type OrderDetail struct {
	// Unique identifier for each pizza placed within each order
	// (pizzas of the same type and size are kept in the same row, and the quantity increases)
	Id int

	// Foreign key that ties the details in each order to the order itself
	OrderId int

	// Foreign key that ties the pizza ordered to its details, like size and price
	PizzaId string

	// Quantity ordered for each pizza of the same type and size
	Quantity int
}

func PrintOrderDetails(orderDetails []OrderDetail) {
	fmt.Println("[Order details]")
	fmt.Println(`"ID", "Order ID", "Pizza ID", "Quantity"`)
	for _, v := range orderDetails {
		fmt.Printf("%d, %d, %q, %d\n", v.Id, v.OrderId, v.PizzaId, v.Quantity)
	}
}

func QueryPrintOrderDetails(db PlatosPizzaDatabase) error {
	records, e := db.HeadOrderDetails()

	if e != nil {
		return trackable.Wrap(e, "Quering head of order_details")
	}

	PrintOrderDetails(records)
	fmt.Println("...")

	return nil
}
