package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

var (
	ErrExecuting = err.Track("Failed to execute workflow")
)

func Execute() error {
	db := database.CreateInMemoryDatabase()

	if e := insertData(db); e != nil {
		return ErrExecuting.TraceWrap(e, "Inserting all data into in-memory database")
	}

	database.Print(db) // Temp

	return nil
}

func insertData(db database.PlatosPizzaDatabase) error {
	e := database.InsertMetadata(db, "../data/data_dictionary.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert metadata")
	}

	e = database.InsertOrders(db, "../data/orders.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert orders")
	}

	e = database.InsertOrderDetails(db, "../data/order_details.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert order details")
	}

	e = database.InsertPizzas(db, "../data/pizzas.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert pizzas")
	}

	e = database.InsertPizzaTypes(db, "../data/pizza_types.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert pizza types")
	}

	return nil
}
