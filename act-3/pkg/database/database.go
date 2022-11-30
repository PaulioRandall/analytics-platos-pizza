package database

import (
	"fmt"

	"github.com/PaulioRandall/trackable"
)

const (
	QueryHeadMax   = 8
	DatetimeFormat = "2006-01-02 15:04:05"
)

var (
	ErrCreating  = trackable.Track("Failed to create database or tables")
	ErrPreparing = trackable.Track("Failed to prepare database query or statement")
	ErrInserting = trackable.Track("Failed to execute data insert into database")
	ErrQuerying  = trackable.Track("Failed to execute query on database")
	ErrParsing   = trackable.Track("Failed to read or parse database results")
	ErrPrinting  = trackable.Track("Failed to print rows from database")
	ErrClosed    = trackable.Track("Can't execute requests on a closed database")
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
	if e := QueryPrintMetadata(db); e != nil {
		return ErrPrinting.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintOrders(db); e != nil {
		return ErrPrinting.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintOrderDetails(db); e != nil {
		return ErrPrinting.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintPizzas(db); e != nil {
		return ErrPrinting.Wrap(e)
	}

	fmt.Println()

	if e := QueryPrintPizzaTypes(db); e != nil {
		return ErrPrinting.Wrap(e)
	}

	return nil
}
