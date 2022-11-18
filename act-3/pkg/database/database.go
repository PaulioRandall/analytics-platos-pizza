package database

import (
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/models"
)

// PlatosPizzaDatabase represents an interface to a database of orders, pizzas,
// and information useful for analysing Plato's Pizzeria customer buying
// habits.
type PlatosPizzaDatabase interface {
	QueryAllDataDicts() ([]models.DataDictionary, error)
	InsertDataDictEntry(entry models.DataDictionary) error
}
