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
