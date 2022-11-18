package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

func insertData(db database.PlatosPizzaDatabase) error {
	return err.ErrTodo.Track("Inserting data")
}
