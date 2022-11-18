package workflow

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

var (
	ErrExecuting = err.NewTrackable("Failed to execute workflow")
)

func Execute() error {
	db := database.CreateInMemoryDatabase()

	if e := insertData(db); e != nil {
		return ErrExecuting.Wrap(e)
	}

	// Temp
	if dataDicts, e := db.QueryAllDataDicts(); e != nil {
		return ErrExecuting.Wrap(e)
	} else {
		models.PrintDataDictionary(dataDicts)
	}

	return nil
}
