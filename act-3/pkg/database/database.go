package database

// PlatosPizzaDatabase represents an interface to a database of orders, pizzas,
// and information useful for analysing Plato's Pizzeria customer buying
// habits.
type PlatosPizzaDatabase interface {
	InsertMetadata(entry MetadataEntry) error
	QueryAllMetadata() ([]MetadataEntry, error)

	InsertOrder(order Order) error
	QueryHeadOrders() ([]Order, error)
}
