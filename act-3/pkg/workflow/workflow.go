package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

var (
	ErrExecuting = err.Track("Failed to execute workflow")
)

func Execute() error {
	db := database.CreateInMemoryDatabase()

	if e := insertData(db); e != nil {
		return ErrExecuting.TraceWrap(e, "Inserting all data into in-memory database")
	}

	// Temp
	if dataDicts, e := db.QueryAllDataDicts(); e != nil {
		return ErrExecuting.TraceWrap(e, "Quering all data dictionary entries")
	} else {
		models.PrintDataDictionary(dataDicts)
	}

	return nil
}
