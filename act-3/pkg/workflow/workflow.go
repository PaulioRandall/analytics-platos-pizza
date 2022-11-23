package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

var (
	ErrExecuting = err.Track("Failed to execute workflow")
)

func Execute() error {
	db := database.NewInMemoryDatabase()

	if e := insertCSVData(db); e != nil {
		return ErrExecuting.TraceWrap(e, "Inserting all CSV data into database")
	}

	database.Print(db) // Temp

	return nil
}

func insertCSVData(db database.PlatosPizzaDatabase) error {
	e := database.InsertMetadataFromCSV(db, "../data/data_dictionary.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert metadata")
	}

	e = database.InsertOrdersFromCSV(db, "../data/orders.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert orders")
	}

	e = database.InsertOrderDetailsFromCSV(db, "../data/order_details.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert order details")
	}

	e = database.InsertPizzasFromCSV(db, "../data/pizzas.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert pizzas")
	}

	e = database.InsertPizzaTypesFromCSV(db, "../data/pizza_types.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert pizza types")
	}

	return nil
}
