package workflow

import (
	"github.com/PaulioRandall/trackable"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/scene-2/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/scene-2/database/sqlite"
)

var (
	ErrExeWorkflow = trackable.Track("Failed to execute workflow")
)

func Execute() error {
	db, e := sqlite.OpenDatabase("./bin/platos-pizza.sqlite")
	if e != nil {
		return ErrExeWorkflow.Wrap(e)
	}
	defer db.Close()

	if e := insertCSVData(db); e != nil {
		return ErrExeWorkflow.Wrap(e)
	}

	// Temp
	if e := database.Print(db); e != nil {
		return ErrExeWorkflow.Wrap(e)
	}

	return nil
}

func insertCSVData(db database.PlatosPizzaDatabase) error {
	file := "../data/data_dictionary.csv"
	e := database.InsertMetadataFromCSV(db, file)
	if e != nil {
		return trackable.Wrap(e, "Failed to insert metadata from %q", file)
	}

	file = "../data/orders.csv"
	e = database.InsertOrdersFromCSV(db, file)
	if e != nil {
		return trackable.Wrap(e, "Failed to insert orders from %q", file)
	}

	file = "../data/order_details.csv"
	e = database.InsertOrderDetailsFromCSV(db, file)
	if e != nil {
		return trackable.Wrap(e, "Failed to insert order details from %q", file)
	}

	file = "../data/pizzas.csv"
	e = database.InsertPizzasFromCSV(db, file)
	if e != nil {
		return trackable.Wrap(e, "Failed to insert pizzas from %q", file)
	}

	file = "../data/pizza_types.csv"
	e = database.InsertPizzaTypesFromCSV(db, file)
	if e != nil {
		return trackable.Wrap(e, "Failed to insert pizza types from %q", file)
	}

	return nil
}
