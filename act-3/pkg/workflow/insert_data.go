package workflow

import (
	"encoding/csv"
	"os"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

func insertData(db database.PlatosPizzaDatabase) error {
	e := insertMetadata(db, "../data/data_dictionary.csv")
	if e != nil {
		return e
	}

	return nil
}

func insertMetadata(db database.PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return err.Wrap(e, "Failed to read metadata %q", filename)
	}

	for i, v := range records {
		m := parseMetadataEntry(v)
		e := db.InsertMetadata(m)
		if e != nil {
			return err.Wrap(e, "Failed to insert metadata record at line %d", i+1)
		}
	}

	return nil
}

func parseMetadataEntry(record []string) database.MetadataEntry {
	return database.MetadataEntry{
		Table:       record[0],
		Field:       record[1],
		Description: record[2],
	}
}

func readCSV(filename string) ([][]string, error) {
	readErr := err.Track("Error reading CSV file %q", filename)

	f, e := os.Open(filename)
	if e != nil {
		return nil, readErr.Wrap(e)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, e := r.ReadAll()
	records = records[1:] // Remove header

	if e != nil {
		return nil, readErr.Wrap(e)
	}

	return records, nil
}
