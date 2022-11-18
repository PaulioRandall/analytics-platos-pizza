package models

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

// OrderDetails represents a specific pizza type order, one or more pizzas,
// within an order
type OrderDetails struct {
	// Unique identifier for each pizza placed within each order
	// (pizzas of the same type and size are kept in the same row, and the quantity increases)
	Id int

	// Foreign key that ties the details in each order to the order itself
	OrderId int

	// Foreign key that ties the pizza ordered to its details, like size and price
	PizzaId int

	// Quantity ordered for each pizza of the same type and size
	Quantity int
}

type Pizzas struct {
	// Unique identifier for each pizza (constituted by its type and size)
	Id int

	// Foreign key that ties each pizza to its broader pizza type
	PizzaTypeId int

	// Size of the pizza (Small, Medium, Large, X Large, or XX Large)
	Size string

	// Price of the pizza in USD
	Price float64
}

type PizzaTypes struct {
	// Unique identifier for each pizza type
	Id int

	// Name of the pizza as shown in the menu
	Name string

	// Category that the pizza fall under in the menu
	// (Classic, Chicken, Supreme, or Veggie)
	Category string

	// Comma-delimited ingredients used in the pizza as shown in the menu
	// (they all include Mozzarella Cheese, even if not specified; and they all
	// include Tomato Sauce, unless another sauce is specified)
	Ingredients string
}
