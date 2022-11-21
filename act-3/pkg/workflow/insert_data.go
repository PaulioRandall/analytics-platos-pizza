package workflow

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

func insertData(db database.PlatosPizzaDatabase) error {
	e := insertMetadata(db, "../data/data_dictionary.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert metadata")
	}

	e = insertOrders(db, "../data/orders.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert orders")
	}

	e = insertOrderDetails(db, "../data/order_details.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert order details")
	}

	e = insertPizzas(db, "../data/pizzas.csv")
	if e != nil {
		return err.Wrap(e, "Failed to insert pizzas")
	}

	return nil
}

func insertMetadata(db database.PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return err.Wrap(e, "Failed to read metadata %q", filename)
	}

	for i, record := range records {
		entry := database.MetadataEntry{
			Table:       record[0],
			Field:       record[1],
			Description: record[2],
		}

		e := db.InsertMetadata(entry)
		if e != nil {
			return err.Wrap(e, "Failed to insert metadata record at line %d", i+1)
		}
	}

	return nil
}

func insertOrders(db database.PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return err.Wrap(e, "Failed to read orders %q", filename)
	}

	for i, record := range records {
		id, e := strconv.Atoi(record[0])
		if e != nil {
			return err.Wrap(e, "Bad order ID discovered")
		}

		datetime, e := time.Parse(time.RFC3339, record[1]+"T"+record[2]+"Z")
		if e != nil {
			return err.Wrap(e, "Bad order date or time discovered")
		}

		order := database.Order{
			Id:       id,
			Datetime: datetime,
		}

		if e = db.InsertOrder(order); e != nil {
			return err.Wrap(e, "Failed to insert order record at line %d", i+1)
		}
	}

	return nil
}

func insertOrderDetails(db database.PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return err.Wrap(e, "Failed to read order details %q", filename)
	}

	for i, record := range records {
		id, e := strconv.Atoi(record[0])
		if e != nil {
			return err.Wrap(e, "Bad order details ID discovered")
		}

		orderId, e := strconv.Atoi(record[1])
		if e != nil {
			return err.Wrap(e, "Bad order ID discovered")
		}

		quantity, e := strconv.Atoi(record[3])
		if e != nil {
			return err.Wrap(e, "Bad quantity value discovered")
		}

		orderDetail := database.OrderDetail{
			Id:       id,
			OrderId:  orderId,
			PizzaId:  record[2],
			Quantity: quantity,
		}

		if e = db.InsertOrderDetail(orderDetail); e != nil {
			return err.Wrap(e, "Failed to insert order detail at line %d", i+1)
		}
	}

	return nil
}

func insertPizzas(db database.PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return err.Wrap(e, "Failed to read pizzas %q", filename)
	}

	for i, record := range records {
		price, e := strconv.ParseFloat(record[3], 64)
		if e != nil {
			return err.Wrap(e, "Bad price value discovered")
		}

		pizza := database.Pizza{
			Id:          record[0],
			PizzaTypeId: record[1],
			Size:        record[2],
			Price:       price,
		}

		if e = db.InsertPizza(pizza); e != nil {
			return err.Wrap(e, "Failed to insert pizza at line %d", i+1)
		}
	}

	return nil
}

func readCSV(filename string) ([][]string, error) {
	readErr := err.Track("Error reading CSV file %q", filename)

	f, e := os.Open(filename)
	if e != nil {
		return nil, readErr.Wrap(e)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, e := r.ReadAll()
	records = records[1:] // Remove header

	if e != nil {
		return nil, readErr.Wrap(e)
	}

	return records, nil
}
