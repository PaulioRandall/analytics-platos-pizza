package database

import (
	"fmt"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

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

func QueryPrintMetadata(db PlatosPizzaDatabase) error {
	records, e := db.QueryAllMetadata()

	if e != nil {
		return err.Wrap(e, "Quering all metadata")
	}

	PrintMetadata(records)
	return nil
}

func QueryPrintOrders(db PlatosPizzaDatabase) error {
	records, e := db.QueryHeadOrders()

	if e != nil {
		return err.Wrap(e, "Quering all metadata")
	}

	PrintOrders(records)
	return nil
}
