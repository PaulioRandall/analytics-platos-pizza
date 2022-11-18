package database

import (
	//"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

type inMemory struct {
	dataDictionaries []models.DataDictionary
	orders           []models.Orders
	order_details    []models.OrderDetails
	pizzas           []models.Pizzas
	pizza_types      []models.PizzaTypes
}

func CreateInMemoryDatabase() *inMemory {
	return &inMemory{}
}

func (db *inMemory) InsertDataDictEntry(entry models.DataDictionary) error {
	db.dataDictionaries = append(db.dataDictionaries, entry)
	return nil
}

func (db *inMemory) QueryAllDataDicts() ([]models.DataDictionary, error) {
	return db.dataDictionaries, nil
}
