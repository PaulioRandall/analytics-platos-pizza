package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

var (
	ErrWorkflow = err.Track("Failed to execute workflow")
)

func Execute() error {
	db, e := database.OpenSQLiteDatabase("./build/platos-pizza.sqlite")
	if e != nil {
		return ErrWorkflow.TraceWrap(e, "Failed to open database")
	}
	defer db.Close()

	if e := insertCSVData(db); e != nil {
		return ErrWorkflow.TraceWrap(e, "Failed to insert CSV data into database")
	}

	// Temp
	if e := database.Print(db); e != nil {
		return ErrWorkflow.TraceWrap(e, "Failed to print database")
	}

	return nil
}

func insertCSVData(db database.PlatosPizzaDatabase) error {
	file := "../data/data_dictionary.csv"
	e := database.InsertMetadataFromCSV(db, file)
	if e != nil {
		return err.Wrap(e, "Failed to insert metadata from %q", file)
	}

	file = "../data/orders.csv"
	e = database.InsertOrdersFromCSV(db, file)
	if e != nil {
		return err.Wrap(e, "Failed to insert orders from %q", file)
	}

	file = "../data/order_details.csv"
	e = database.InsertOrderDetailsFromCSV(db, file)
	if e != nil {
		return err.Wrap(e, "Failed to insert order details from %q", file)
	}

	file = "../data/pizzas.csv"
	e = database.InsertPizzasFromCSV(db, file)
	if e != nil {
		return err.Wrap(e, "Failed to insert pizzas from %q", file)
	}

	file = "../data/pizza_types.csv"
	e = database.InsertPizzaTypesFromCSV(db, file)
	if e != nil {
		return err.Wrap(e, "Failed to insert pizza types from %q", file)
	}

	return nil
}
