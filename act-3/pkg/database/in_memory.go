package database

import (
	"github.com/PaulioRandall/trackable-go"
)

var ErrInMemory = trackable.Track("In-memory database error")

type inMemory struct {
	closed       bool
	metadata     []MetadataEntry
	orders       []Order
	orderDetails []OrderDetail
	pizzas       []Pizza
	pizzaTypes   []PizzaType
}

func OpenInMemoryDatabase() *inMemory {
	return &inMemory{}
}

func (db *inMemory) InsertMetadata(entries ...MetadataEntry) error {
	return inMemoryInsert(db, func() {
		for _, v := range entries {
			db.metadata = append(db.metadata, v)
		}
	})
}

func (db *inMemory) InsertOrders(orders ...Order) error {
	return inMemoryInsert(db, func() {
		for _, v := range orders {
			db.orders = append(db.orders, v)
		}
	})
}

func (db *inMemory) InsertOrderDetails(orderDetails ...OrderDetail) error {
	return inMemoryInsert(db, func() {
		for _, v := range orderDetails {
			db.orderDetails = append(db.orderDetails, v)
		}
	})
}

func (db *inMemory) InsertPizzas(pizzas ...Pizza) error {
	return inMemoryInsert(db, func() {
		for _, v := range pizzas {
			db.pizzas = append(db.pizzas, v)
		}
	})
}

func (db *inMemory) InsertPizzaTypes(pizzaTypes ...PizzaType) error {
	return inMemoryInsert(db, func() {
		for _, v := range pizzaTypes {
			db.pizzaTypes = append(db.pizzaTypes, v)
		}
	})
}

func (db *inMemory) AllMetadata() ([]MetadataEntry, error) {
	return inMemoryExecute(db, func() ([]MetadataEntry, error) {
		return db.metadata[:], nil
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

	db.metadata = nil
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
		e = ErrQuerying.Wrap(e)
		e = ErrInMemory.Wrap(e)
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
