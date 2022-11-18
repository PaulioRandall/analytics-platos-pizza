package workflow

import (
	"encoding/csv"
	"os"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

var (
	ErrDataInsert = err.NewTrackable("Failed to insert data")
)

func insertData(db database.PlatosPizzaDatabase) error {
	records, e := readCSV("../data/data_dictionary.csv")
	if e != nil {
		return ErrDataInsert.Wrap(e)
	}

	models := parseDataDictionary(records)
	for _, m := range models {
		db.InsertDataDictEntry(m)
	}

	return nil
}

func readCSV(filename string) ([][]string, error) {
	f, e := os.Open(filename)
	if e != nil {
		return nil, err.NewTrackable("Error reading CSV file %q", filename).Wrap(e)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, e := r.ReadAll()
	records = records[1:] // Remove header

	if e != nil {
		return nil, err.NewTrackable("Error reading CSV file %q", filename).Wrap(e)
	}

	return records, nil
}

func parseDataDictionary(data [][]string) []models.DataDictionary {
	results := make([]models.DataDictionary, len(data))

	for i, v := range data {
		m := parseDictionaryEntry(v)
		results[i] = m
	}

	return results
}

func parseDictionaryEntry(record []string) models.DataDictionary {
	return models.DataDictionary{
		Table:       record[0],
		Field:       record[1],
		Description: record[2],
	}
}
