package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

func Execute() error {

	db := database.CreateInMemoryDatabase()

	if e := insertData(db); e != nil {
		return e
	}

	return err.ErrTodo.Track("Execute workflow")
}
