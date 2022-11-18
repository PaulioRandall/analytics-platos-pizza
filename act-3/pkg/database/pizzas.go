package database

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
