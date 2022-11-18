package database

import (
	//"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

type inMemory struct {
	dataDictionaries []MetadataEntry
	orders           []models.Orders
	order_details    []models.OrderDetails
	pizzas           []models.Pizzas
	pizza_types      []models.PizzaTypes
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
