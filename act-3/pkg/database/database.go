package database

import (
	"fmt"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

// PlatosPizzaDatabase represents an interface to a database of orders, pizzas,
// and information useful for analysing Plato's Pizzeria customer buying
// habits.
type PlatosPizzaDatabase interface {
	InsertMetadata(MetadataEntry) error
	InsertOrder(Order) error
	InsertOrderDetail(OrderDetail) error
	InsertPizza(Pizza) error
	InsertPizzaType(PizzaType) error

	AllMetadata() ([]MetadataEntry, error)
	HeadOrders() ([]Order, error)
	HeadOrderDetails() ([]OrderDetail, error)
	HeadPizzas() ([]Pizza, error)
	HeadPizzaTypes() ([]PizzaType, error)
}

func Print(db PlatosPizzaDatabase) error {
	printErr := err.Track("Printing database")

	if e := QueryPrintMetadata(db); e != nil {
		return printErr.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintOrders(db); e != nil {
		return printErr.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintOrderDetails(db); e != nil {
		return printErr.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintPizzas(db); e != nil {
		return printErr.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintPizzaTypes(db); e != nil {
		return printErr.Wrap(e)
	}

	return nil
}
