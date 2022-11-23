package database

const queryHeadMax = 8

type inMemory struct {
	dataDictionaries []MetadataEntry
	orders           []Order
	orderDetails     []OrderDetail
	pizzas           []Pizza
	pizzaTypes       []PizzaType
}

func NewInMemoryDatabase() *inMemory {
	return &inMemory{}
}

func (db *inMemory) InsertMetadata(entry MetadataEntry) error {
	db.dataDictionaries = append(db.dataDictionaries, entry)
	return nil
}

func (db *inMemory) InsertOrder(order Order) error {
	db.orders = append(db.orders, order)
	return nil
}

func (db *inMemory) InsertOrderDetail(orderDetail OrderDetail) error {
	db.orderDetails = append(db.orderDetails, orderDetail)
	return nil
}

func (db *inMemory) InsertPizza(pizza Pizza) error {
	db.pizzas = append(db.pizzas, pizza)
	return nil
}

func (db *inMemory) InsertPizzaType(pizzaType PizzaType) error {
	db.pizzaTypes = append(db.pizzaTypes, pizzaType)
	return nil
}

func (db *inMemory) AllMetadata() ([]MetadataEntry, error) {
	return db.dataDictionaries[:], nil
}

func (db *inMemory) HeadOrders() ([]Order, error) {
	return head(db.orders)
}

func (db *inMemory) HeadOrderDetails() ([]OrderDetail, error) {
	return head(db.orderDetails)
}

func (db *inMemory) HeadPizzas() ([]Pizza, error) {
	return head(db.pizzas)
}

func (db *inMemory) HeadPizzaTypes() ([]PizzaType, error) {
	return head(db.pizzaTypes)
}

func head[T any](items []T) ([]T, error) {
	if len(items) < queryHeadMax {
		return items[:], nil
	}
	return items[0:queryHeadMax], nil
}
