package database

type PizzaType struct {
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
