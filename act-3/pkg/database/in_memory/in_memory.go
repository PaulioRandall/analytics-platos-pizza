package in_memory

import (
	"github.com/PaulioRandall/trackable"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
)

const QueryHeadMax = database.QueryHeadMax

var ErrInMemory = trackable.Interface("In-memory database error")

type query[T any] func() ([]T, error)

type inMemory struct {
	closed       bool
	metadata     []database.MetadataEntry
	orders       []database.Order
	orderDetails []database.OrderDetail
	pizzas       []database.Pizza
	pizzaTypes   []database.PizzaType
}

func OpenDatabase() *inMemory {
	return &inMemory{}
}

func (db *inMemory) InsertMetadata(entries ...database.MetadataEntry) error {
	return inMemoryInsert(db, func() {
		for _, v := range entries {
			db.metadata = append(db.metadata, v)
		}
	})
}

func (db *inMemory) InsertOrders(orders ...database.Order) error {
	return inMemoryInsert(db, func() {
		for _, v := range orders {
			db.orders = append(db.orders, v)
		}
	})
}

func (db *inMemory) InsertOrderDetails(orderDetails ...database.OrderDetail) error {
	return inMemoryInsert(db, func() {
		for _, v := range orderDetails {
			db.orderDetails = append(db.orderDetails, v)
		}
	})
}

func (db *inMemory) InsertPizzas(pizzas ...database.Pizza) error {
	return inMemoryInsert(db, func() {
		for _, v := range pizzas {
			db.pizzas = append(db.pizzas, v)
		}
	})
}

func (db *inMemory) InsertPizzaTypes(pizzaTypes ...database.PizzaType) error {
	return inMemoryInsert(db, func() {
		for _, v := range pizzaTypes {
			db.pizzaTypes = append(db.pizzaTypes, v)
		}
	})
}

func (db *inMemory) AllMetadata() ([]database.MetadataEntry, error) {
	return inMemoryExecute(db, func() ([]database.MetadataEntry, error) {
		return db.metadata, nil
	})
}

func (db *inMemory) HeadOrders() ([]database.Order, error) {
	return inMemoryHead(db, db.orders)
}

func (db *inMemory) HeadOrderDetails() ([]database.OrderDetail, error) {
	return inMemoryHead(db, db.orderDetails)
}

func (db *inMemory) HeadPizzas() ([]database.Pizza, error) {
	return inMemoryHead(db, db.pizzas)
}

func (db *inMemory) HeadPizzaTypes() ([]database.PizzaType, error) {
	return inMemoryHead(db, db.pizzaTypes)
}

func (db *inMemory) Close() {
	db.closed = true

	db.metadata = nil
	db.orders = nil
	db.orderDetails = nil
	db.pizzas = nil
	db.pizzaTypes = nil
}

func inMemoryHead[T any](db *inMemory, items []T) ([]T, error) {
	return inMemoryExecute(db, func() ([]T, error) {
		if len(items) < QueryHeadMax {
			return items, nil
		}
		return items[0:QueryHeadMax], nil
	})
}

func inMemoryExecute[T any](db *inMemory, q query[T]) ([]T, error) {
	if db.closed {
		return nil, ErrInMemory.Wrap(database.ErrClosed)
	}

	result, e := q()
	if e != nil {
		e = database.ErrQuerying.Wrap(e)
		e = ErrInMemory.Wrap(e)
	}

	return result, e
}

func inMemoryInsert(db *inMemory, f func()) error {
	if db.closed {
		return ErrInMemory.Wrap(database.ErrClosed)
	}

	f()
	return nil
}
