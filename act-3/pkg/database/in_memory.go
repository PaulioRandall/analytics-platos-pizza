package database

import (
//"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

type inMemory struct {
	dataDictionaries []MetadataEntry
	orders           []Order
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

func (db *inMemory) InsertOrder(order Order) error {
	db.orders = append(db.orders, order)
	return nil
}

func (db *inMemory) QueryHeadOrders() ([]Order, error) {
	const maxLen = 8

	if len(db.orders) < maxLen {
		return db.orders[:], nil
	}

	return db.orders[0:8], nil
}
