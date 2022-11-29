package database

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/PaulioRandall/trackable-go"
)

var (
	ErrCSVFile = trackable.Track("Error handling CSV file")
)

func InsertMetadataFromCSV(db PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return trackable.Wrap(e, "Failed to read metadata %q", filename)
	}

	for i, record := range records {
		entry := MetadataEntry{
			Table:       record[0],
			Field:       record[1],
			Description: record[2],
		}

		e := db.InsertMetadata(entry)
		if e != nil {
			return trackable.Wrap(e,
				"Failed to insert metadata record at line %d", lineNumber(i),
			)
		}
	}

	return nil
}

func InsertOrdersFromCSV(db PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return trackable.Wrap(e, "Failed to read orders %q", filename)
	}

	orders := make([]Order, len(records))
	for i, record := range records {
		id, e := strconv.Atoi(record[0])
		if e != nil {
			return trackable.Wrap(e, "Bad order ID discovered")
		}

		datetime, e := time.Parse(DatetimeFormat, record[1]+" "+record[2])
		if e != nil {
			return trackable.Wrap(e, "Bad order date or time discovered")
		}

		orders[i] = Order{
			Id:       id,
			Datetime: datetime,
		}
	}

	if e = db.InsertOrders(orders...); e != nil {
		return e
	}

	return nil
}

func InsertOrderDetailsFromCSV(db PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return trackable.Wrap(e, "Failed to read order details %q", filename)
	}

	for i, record := range records {
		id, e := strconv.Atoi(record[0])
		if e != nil {
			return trackable.Wrap(e, "Bad order details ID discovered")
		}

		orderId, e := strconv.Atoi(record[1])
		if e != nil {
			return trackable.Wrap(e, "Bad order ID discovered")
		}

		quantity, e := strconv.Atoi(record[3])
		if e != nil {
			return trackable.Wrap(e, "Bad quantity value discovered")
		}

		orderDetail := OrderDetail{
			Id:       id,
			OrderId:  orderId,
			PizzaId:  record[2],
			Quantity: quantity,
		}

		if e = db.InsertOrderDetails(orderDetail); e != nil {
			return trackable.Wrap(e,
				"Failed to insert order detail at line %d", lineNumber(i),
			)
		}
	}

	return nil
}

func InsertPizzasFromCSV(db PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return trackable.Wrap(e, "Failed to read pizzas %q", filename)
	}

	for i, record := range records {
		price, e := strconv.ParseFloat(record[3], 64)
		if e != nil {
			return trackable.Wrap(e, "Bad price value discovered")
		}

		pizza := Pizza{
			Id:          record[0],
			PizzaTypeId: record[1],
			Size:        record[2],
			Price:       price,
		}

		if e = db.InsertPizzas(pizza); e != nil {
			return trackable.Wrap(e, "Failed to insert pizza at line %d", lineNumber(i))
		}
	}

	return nil
}

func InsertPizzaTypesFromCSV(db PlatosPizzaDatabase, filename string) error {
	records, e := readCSV(filename)
	if e != nil {
		return trackable.Wrap(e, "Failed to read pizza types %q", filename)
	}

	for i, record := range records {
		pizzaType := PizzaType{
			Id:          record[0],
			Name:        record[1],
			Category:    record[2],
			Ingredients: record[3],
		}

		if e = db.InsertPizzaTypes(pizzaType); e != nil {
			return trackable.Wrap(e,
				"Failed to insert pizza type at line %d", lineNumber(i),
			)
		}
	}

	return nil
}

func lineNumber(i int) int {
	i++ // Convert from index to count
	i++ // Skip the header
	return i
}

func readCSV(filename string) ([][]string, error) {
	f, e := os.Open(filename)
	if e != nil {
		return nil, ErrCSVFile.BecauseOf(e, "Could not open file %q", filename)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, e := r.ReadAll()
	records = records[1:] // Remove header

	if e != nil {
		return nil, ErrCSVFile.BecauseOf(e, "Could not read file %q", filename)
	}

	return records, nil
}
