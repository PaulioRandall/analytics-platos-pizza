package workflow

import (
	"encoding/csv"
	"os"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

var (
	ErrDataInsert = err.Track("Failed to insert data")
)

func insertData(db database.PlatosPizzaDatabase) error {
	records, e := readCSV("../data/data_dictionary.csv")
	if e != nil {
		return ErrDataInsert.Wrap(e)
	}

	models := parseMetadata(records)
	for _, m := range models {
		db.InsertMetadata(m)
	}

	return nil
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

func parseMetadata(data [][]string) []database.MetadataEntry {
	results := make([]database.MetadataEntry, len(data))

	for i, v := range data {
		m := parseMetadataEntry(v)
		results[i] = m
	}

	return results
}

func parseMetadataEntry(record []string) database.MetadataEntry {
	return database.MetadataEntry{
		Table:       record[0],
		Field:       record[1],
		Description: record[2],
	}
}
