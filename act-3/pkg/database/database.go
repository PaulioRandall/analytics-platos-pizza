package database

import (
	"fmt"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

// PlatosPizzaDatabase represents an interface to a database of orders, pizzas,
// and information useful for analysing Plato's Pizzeria customer buying
// habits.
type PlatosPizzaDatabase interface {
	InsertMetadata(entry MetadataEntry) error
	QueryAllMetadata() ([]MetadataEntry, error)

	InsertOrder(order Order) error
	QueryHeadOrders() ([]Order, error)
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

	return nil
}
