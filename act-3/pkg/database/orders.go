package database

import (
	"time"
)

// Orders represents an order of pizzas, one or many pizzas per order
type Orders struct {
	// Unique identifier for each order placed by a table
	Id int

	// Date & time the order was placed
	// (entered into the system prior to cooking & serving)
	Datetime time.Time
}
