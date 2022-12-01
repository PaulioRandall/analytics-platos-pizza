package database

import (
	"database/sql"
	"time"
)

func (db *sqliteDB) AllMetadata() ([]MetadataEntry, error) {
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
		return nil, ErrQuerying.Wrap(e)
	}
	defer rows.Close()

	return scanMetadataRows(rows)
}

func scanMetadataRows(rows *sql.Rows) ([]MetadataEntry, error) {
	var results []MetadataEntry

	for rows.Next() {
		var entry MetadataEntry

		e := rows.Scan(&entry.Table, &entry.Field, &entry.Description)
		if e != nil {
			return nil, ErrParsing.Wrap(e)
		}

		results = append(results, entry)
	}

	return results, nil
}

func (db *sqliteDB) HeadOrders() ([]Order, error) {
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
		return nil, ErrQuerying.Wrap(e)
	}
	defer rows.Close()

	return scanOrderRows(rows)
}

func scanOrderRows(rows *sql.Rows) ([]Order, error) {
	var results []Order

	for rows.Next() {
		var order Order
		var datetimeStr string

		e := rows.Scan(&order.Id, &datetimeStr)
		if e != nil {
			return nil, ErrParsing.Wrap(e)
		}

		order.Datetime, e = time.Parse(DatetimeFormat, datetimeStr)
		if e != nil {
			return nil, ErrParsing.Wrap(e)
		}

		results = append(results, order)
	}

	return results, nil
}

func (db *sqliteDB) HeadOrderDetails() ([]OrderDetail, error) {
	return nil, nil
}

func (db *sqliteDB) HeadPizzas() ([]Pizza, error) {
	return nil, nil
}

func (db *sqliteDB) HeadPizzaTypes() ([]PizzaType, error) {
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
		return nil, ErrQuerying.Wrap(e)
	}
	defer rows.Close()

	return scanPizzaTypeRows(rows)
}

func scanPizzaTypeRows(rows *sql.Rows) ([]PizzaType, error) {
	var results []PizzaType

	for rows.Next() {
		var pizzaType PizzaType

		e := rows.Scan(
			&pizzaType.Id,
			&pizzaType.Name,
			&pizzaType.Category,
			&pizzaType.Ingredients,
		)

		if e != nil {
			return nil, ErrParsing.BecauseOf(e, "Row scanning failed")
		}

		results = append(results, pizzaType)
	}

	return results, nil
}
