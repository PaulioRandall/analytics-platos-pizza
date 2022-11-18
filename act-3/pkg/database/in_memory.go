package database

import (
//"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

const queryHeadMax = 16

type inMemory struct {
	dataDictionaries []MetadataEntry
	orders           []Order
	orderDetails     []OrderDetail
	pizzas           []Pizzas
	pizzaTypes       []PizzaTypes
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
	if len(db.orders) < queryHeadMax {
		return db.orders[:], nil
	}

	return db.orders[0:queryHeadMax], nil
}

func (db *inMemory) InsertOrderDetail(orderDetail OrderDetail) error {
	db.orderDetails = append(db.orderDetails, orderDetail)
	return nil
}

func (db *inMemory) QueryHeadOrderDetails() ([]OrderDetail, error) {
	if len(db.orderDetails) < queryHeadMax {
		return db.orderDetails[:], nil
	}

	return db.orderDetails[0:queryHeadMax], nil
}
