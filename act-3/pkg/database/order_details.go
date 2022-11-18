package database

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
