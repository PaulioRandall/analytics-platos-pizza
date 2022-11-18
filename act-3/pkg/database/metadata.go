package database

import (
	"fmt"
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
