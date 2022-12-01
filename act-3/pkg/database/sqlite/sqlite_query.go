package sqlite

import (
	"database/sql"
	"time"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
)

const QueryHeadMax = database.QueryHeadMax

func (db *sqliteDB) AllMetadata() ([]database.MetadataEntry, error) {
	sql := joinLines(
		`SELECT`,
		`	table_name,`,
		`	field_name,`,
		`	description`,
		`FROM`,
		`	metadata;`,
	)

	rows, e := db.conn.Query(sql)
	if e != nil {
		return nil, database.ErrQuerying.Wrap(e)
	}
	defer rows.Close()

	return scanMetadataRows(rows)
}

func scanMetadataRows(rows *sql.Rows) ([]database.MetadataEntry, error) {
	var results []database.MetadataEntry

	for rows.Next() {
		var entry database.MetadataEntry

		e := rows.Scan(&entry.Table, &entry.Field, &entry.Description)
		if e != nil {
			return nil, database.ErrParsing.Wrap(e)
		}

		results = append(results, entry)
	}

	return results, nil
}

func (db *sqliteDB) HeadOrders() ([]database.Order, error) {
	sql := joinLines(
		`SELECT`,
		`	id,`,
		`	strftime('%Y-%m-%d %H:%M:%S', orders.datetime) AS datetime`,
		`FROM`,
		`	orders`,
		// Could have used the value directly as SQL injection is not possible here
		// But it does mean the SQL driver will handle the type conversion for me
		`LIMIT ?;`,
	)

	rows, e := db.conn.Query(sql, QueryHeadMax)
	if e != nil {
		return nil, database.ErrQuerying.BecauseOf(e, "Querying orders")
	}
	defer rows.Close()

	return scanOrderRows(rows)
}

func scanOrderRows(rows *sql.Rows) ([]database.Order, error) {
	var results []database.Order

	for rows.Next() {
		var order database.Order
		var datetimeStr string

		e := rows.Scan(&order.Id, &datetimeStr)
		if e != nil {
			return nil, database.ErrParsing.Wrap(e)
		}

		order.Datetime, e = time.Parse(database.DatetimeFormat, datetimeStr)
		if e != nil {
			return nil, database.ErrParsing.Wrap(e)
		}

		results = append(results, order)
	}

	return results, nil
}

func (db *sqliteDB) HeadOrderDetails() ([]database.OrderDetail, error) {
	sql := joinLines(
		`SELECT`,
		`	id,`,
		`	order_id,`,
		`	pizza_id,`,
		`	quantity`,
		`FROM`,
		`	order_details`,
		`LIMIT ?;`,
	)

	rows, e := db.conn.Query(sql, QueryHeadMax)
	if e != nil {
		return nil, database.ErrQuerying.BecauseOf(e, "Querying order details")
	}
	defer rows.Close()

	return scanOrderDetailRows(rows)
}

func scanOrderDetailRows(rows *sql.Rows) ([]database.OrderDetail, error) {
	var results []database.OrderDetail

	for rows.Next() {
		var orderDetail database.OrderDetail

		e := rows.Scan(
			&orderDetail.Id,
			&orderDetail.OrderId,
			&orderDetail.PizzaId,
			&orderDetail.Quantity,
		)

		if e != nil {
			return nil, database.ErrParsing.BecauseOf(e, "Row scanning failed")
		}

		results = append(results, orderDetail)
	}

	return results, nil
}

func (db *sqliteDB) HeadPizzas() ([]database.Pizza, error) {
	sql := joinLines(
		`SELECT`,
		`	id,`,
		`	type_id,`,
		`	size,`,
		`	price`,
		`FROM`,
		`	pizzas`,
		`LIMIT ?;`,
	)

	rows, e := db.conn.Query(sql, QueryHeadMax)
	if e != nil {
		return nil, database.ErrQuerying.BecauseOf(e, "Querying pizzas")
	}
	defer rows.Close()

	return scanPizzaRows(rows)
}

func scanPizzaRows(rows *sql.Rows) ([]database.Pizza, error) {
	var results []database.Pizza

	for rows.Next() {
		var pizza database.Pizza

		e := rows.Scan(
			&pizza.Id,
			&pizza.TypeId,
			&pizza.Size,
			&pizza.Price,
		)

		if e != nil {
			return nil, database.ErrParsing.BecauseOf(e, "Row scanning failed")
		}

		results = append(results, pizza)
	}

	return results, nil
}

func (db *sqliteDB) HeadPizzaTypes() ([]database.PizzaType, error) {
	sql := joinLines(
		`SELECT`,
		`	id,`,
		`	name,`,
		`	category,`,
		`	ingredients`,
		`FROM`,
		`	pizza_types`,
		`LIMIT ?;`,
	)

	rows, e := db.conn.Query(sql, QueryHeadMax)
	if e != nil {
		return nil, database.ErrQuerying.BecauseOf(e, "Querying pizza types")
	}
	defer rows.Close()

	return scanPizzaTypeRows(rows)
}

func scanPizzaTypeRows(rows *sql.Rows) ([]database.PizzaType, error) {
	var results []database.PizzaType

	for rows.Next() {
		var pizzaType database.PizzaType

		e := rows.Scan(
			&pizzaType.Id,
			&pizzaType.Name,
			&pizzaType.Category,
			&pizzaType.Ingredients,
		)

		if e != nil {
			return nil, database.ErrParsing.BecauseOf(e, "Row scanning failed")
		}

		results = append(results, pizzaType)
	}

	return results, nil
}