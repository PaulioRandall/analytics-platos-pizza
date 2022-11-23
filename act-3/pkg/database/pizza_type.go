package database

import (
	"fmt"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

type PizzaType struct {
	// Unique identifier for each pizza type
	Id string

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

func PrintPizzaTypes(pizzaTypes []PizzaType) {
	fmt.Println("[Pizza types]")
	fmt.Println(`"ID", "Name", "Category", "Ingredients"`)
	for _, v := range pizzaTypes {
		fmt.Printf("%q, %q, %q, %q\n", v.Id, v.Name, v.Category, v.Ingredients)
	}
}

func QueryPrintPizzaTypes(db PlatosPizzaDatabase) error {
	records, e := db.HeadPizzaTypes()

	if e != nil {
		return err.Wrap(e, "Quering head of pizza types")
	}

	PrintPizzaTypes(records)
	fmt.Println("...")

	return nil
}
