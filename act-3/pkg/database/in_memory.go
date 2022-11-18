package database

import (
//"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

type inMemory struct {
	dataDictionaries []MetadataEntry
	orders           []Orders
	order_details    []OrderDetails
	pizzas           []Pizzas
	pizza_types      []PizzaTypes
}

func CreateInMemoryDatabase() *inMemory {
	return &inMemory{}
}

func (db *inMemory) InsertMetadata(entry MetadataEntry) error {
	db.dataDictionaries = append(db.dataDictionaries, entry)
	return nil
}

func (db *inMemory) QueryAllMetadata() ([]MetadataEntry, error) {
	return db.dataDictionaries, nil
}
