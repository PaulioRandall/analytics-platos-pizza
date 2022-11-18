package database

import (
	"fmt"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

type MetadataEntry struct {
	Table       string
	Field       string
	Description string
}

func PrintMetadata(entries []MetadataEntry) {
	fmt.Println("[Metadata]")
	fmt.Println(`"Table", "Field", "Description"`)
	for _, entry := range entries {
		fmt.Printf("%q, %q, %q\n", entry.Table, entry.Field, entry.Description)
	}
}

func QueryPrintMetadata(db PlatosPizzaDatabase) error {
	records, e := db.QueryAllMetadata()

	if e != nil {
		return err.Wrap(e, "Quering all metadata")
	}

	PrintMetadata(records)
	return nil
}
