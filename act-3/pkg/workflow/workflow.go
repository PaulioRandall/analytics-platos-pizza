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

	// Temp
	if dataDict, e := db.QueryAllMetadata(); e != nil {
		return ErrExecuting.TraceWrap(e, "Quering all data dictionary entries")
	} else {
		database.PrintMetadata(dataDict)
	}

	return nil
}
