package database

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

type inMemory struct {
	dataDictionary models.DataDictionary
	orders         models.Orders
	order_details  models.OrderDetails
	pizzas         models.Pizzas
	pizza_types    models.PizzaTypes
}

func CreateInMemoryDatabase() *inMemory {
	return &inMemory{}
}
