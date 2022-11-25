package database

import (
	"fmt"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

const (
	QueryHeadMax   = 8
	DatetimeFormat = "2006-01-02 15:04:05"
)

var (
	ErrCreateOrUpdate = err.Track("Failed to execute table creation or update")
	ErrPrepare        = err.Track("Failed to prepare statement")
	ErrInsert         = err.Track("Failed to execute data insert")
	ErrQuery          = err.Track("Failed to execute query")
	ErrResult         = err.Track("Failed to read or parse results")
	ErrClosed         = err.Track("Can't execute requests on a closed database")
)

type query[T any] func() ([]T, error)

// PlatosPizzaDatabase represents an interface to a database of orders, pizzas,
// and information useful for analysing Plato's Pizzeria customer buying
// habits.
type PlatosPizzaDatabase interface {
	InsertMetadata(...MetadataEntry) error
	InsertOrders(...Order) error
	InsertOrderDetails(...OrderDetail) error
	InsertPizzas(...Pizza) error
	InsertPizzaTypes(...PizzaType) error

	AllMetadata() ([]MetadataEntry, error)
	HeadOrders() ([]Order, error)
	HeadOrderDetails() ([]OrderDetail, error)
	HeadPizzas() ([]Pizza, error)
	HeadPizzaTypes() ([]PizzaType, error)

	Close() // Panics if error encountered
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
