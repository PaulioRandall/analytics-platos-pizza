package database

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

var ErrInMemory = err.Track("In-memory database error")

type inMemory struct {
	closed           bool
	dataDictionaries []MetadataEntry
	orders           []Order
	orderDetails     []OrderDetail
	pizzas           []Pizza
	pizzaTypes       []PizzaType
}

func OpenInMemoryDatabase() *inMemory {
	return &inMemory{}
}

func (db *inMemory) InsertMetadata(entry MetadataEntry) error {
	return inMemoryInsert(db, func() {
		db.dataDictionaries = append(db.dataDictionaries, entry)
	})
}

func (db *inMemory) InsertOrder(order Order) error {
	return inMemoryInsert(db, func() {
		db.orders = append(db.orders, order)
	})
}

func (db *inMemory) InsertOrderDetail(orderDetail OrderDetail) error {
	return inMemoryInsert(db, func() {
		db.orderDetails = append(db.orderDetails, orderDetail)
	})
}

func (db *inMemory) InsertPizza(pizza Pizza) error {
	return inMemoryInsert(db, func() {
		db.pizzas = append(db.pizzas, pizza)
	})
}

func (db *inMemory) InsertPizzaType(pizzaType PizzaType) error {
	return inMemoryInsert(db, func() {
		db.pizzaTypes = append(db.pizzaTypes, pizzaType)
	})
}

func (db *inMemory) AllMetadata() ([]MetadataEntry, error) {
	return inMemoryExecute(db, func() ([]MetadataEntry, error) {
		return db.dataDictionaries[:], nil
	})
}

func (db *inMemory) HeadOrders() ([]Order, error) {
	return inMemoryExecute(db, func() ([]Order, error) {
		return inMemoryHead(db.orders)
	})
}

func (db *inMemory) HeadOrderDetails() ([]OrderDetail, error) {
	return inMemoryExecute(db, func() ([]OrderDetail, error) {
		return inMemoryHead(db.orderDetails)
	})
}

func (db *inMemory) HeadPizzas() ([]Pizza, error) {
	return inMemoryExecute(db, func() ([]Pizza, error) {
		return inMemoryHead(db.pizzas)
	})
}

func (db *inMemory) HeadPizzaTypes() ([]PizzaType, error) {
	return inMemoryExecute(db, func() ([]PizzaType, error) {
		return inMemoryHead(db.pizzaTypes)
	})
}

func (db *inMemory) Close() {
	db.closed = true

	db.dataDictionaries = nil
	db.orders = nil
	db.orderDetails = nil
	db.pizzas = nil
	db.pizzaTypes = nil
}

func inMemoryExecute[T any](db *inMemory, q query[T]) ([]T, error) {
	if db.closed {
		return nil, ErrInMemory.Wrap(ErrClosed)
	}

	result, e := q()
	if e != nil {
		e = ErrInMemory.Wrap(ErrQuery, e)
	}

	return result, e
}

func inMemoryInsert(db *inMemory, f func()) error {
	if db.closed {
		return ErrInMemory.Wrap(ErrClosed)
	}

	f()
	return nil
}

func inMemoryHead[T any](items []T) ([]T, error) {
	if len(items) < QueryHeadMax {
		return items[:], nil
	}
	return items[0:QueryHeadMax], nil
}
