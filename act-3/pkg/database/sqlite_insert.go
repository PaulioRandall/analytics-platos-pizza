package database

import (
	"strings"
)

// TODO: Convert to using transactions for bulk inserts

type sqlBuilder[T any] func([]T) (sql string, params []any)

const insertBatchSize = 256

func (db *sqliteDB) InsertMetadata(entries ...MetadataEntry) error {
	rowCount := len(entries)
	paramCount := 3

	valuesSQL := buildValuesSQL(rowCount, paramCount)
	sql := joinLines(
		`INSERT INTO metadata (`,
		`	table_name,`,
		`	field_name,`,
		`	description`,
		`) VALUES `+valuesSQL+";",
	)

	var params []any
	for _, v := range entries {
		params = append(params, v.Table, v.Field, v.Description)
	}

	if e := db.insert(sql, params); e != nil {
		return ErrSQLite.Wrap(e)
	}

	return nil
}

func (db *sqliteDB) InsertOrders(orders ...Order) error {
	buildOrdersInsertSQL := func(batch []Order) (sql string, params []any) {
		rowCount := len(batch)
		paramCount := 2

		valuesSQL := buildValuesSQL(rowCount, paramCount)
		sql = joinLines(
			`INSERT INTO orders (`,
			`	id,`,
			`	datetime`,
			`) VALUES `+valuesSQL+";",
		)

		for _, v := range batch {
			params = append(params, v.Id, v.Datetime)
		}

		return sql, params
	}

	return sqlitePartitionedInsert(db, orders, buildOrdersInsertSQL)
}

func (db *sqliteDB) InsertOrderDetails(orderDetails ...OrderDetail) error {
	buildOrderDetailsInsertSQL := func(batch []OrderDetail) (sql string, params []any) {
		rowCount := len(batch)
		paramCount := 4

		valuesSQL := buildValuesSQL(rowCount, paramCount)
		sql = joinLines(
			`INSERT INTO order_details (`,
			`	id,`,
			`	order_id,`,
			`	pizza_id,`,
			`	quantity`,
			`) VALUES `+valuesSQL+";",
		)

		for _, v := range batch {
			params = append(params, v.Id, v.OrderId, v.PizzaId, v.Quantity)
		}

		return sql, params
	}

	return sqlitePartitionedInsert(db, orderDetails, buildOrderDetailsInsertSQL)
}

func (db *sqliteDB) InsertPizzas(pizzas ...Pizza) error {
	buildPizzasInsertSQL := func(batch []Pizza) (sql string, params []any) {
		rowCount := len(batch)
		paramCount := 4

		valuesSQL := buildValuesSQL(rowCount, paramCount)
		sql = joinLines(
			`INSERT INTO pizzas (`,
			`	id,`,
			`	type_id,`,
			`	size,`,
			`	price`,
			`) VALUES `+valuesSQL+";",
		)

		for _, v := range batch {
			params = append(params, v.Id, v.TypeId, v.Size, v.Price)
		}

		return sql, params
	}

	return sqlitePartitionedInsert(db, pizzas, buildPizzasInsertSQL)
}

func (db *sqliteDB) InsertPizzaTypes(pizzaTypes ...PizzaType) error {
	buildPizzaTypesInsertSQL := func(batch []PizzaType) (sql string, params []any) {
		rowCount := len(batch)
		paramCount := 4

		valuesSQL := buildValuesSQL(rowCount, paramCount)
		sql = joinLines(
			`INSERT INTO pizza_types (`,
			`	id,`,
			`	name,`,
			`	category,`,
			`	ingredients`,
			`) VALUES `+valuesSQL+";",
		)

		for _, v := range batch {
			params = append(params, v.Id, v.Name, v.Category, v.Ingredients)
		}

		return sql, params
	}

	return sqlitePartitionedInsert(db, pizzaTypes, buildPizzaTypesInsertSQL)
}

func (db *sqliteDB) insert(sql string, params []any) error {
	stmt, e := db.conn.Prepare(sql)
	if e != nil {
		e = ErrPreparing.Wrap(e)
		return ErrInserting.Wrap(e)
	}
	defer stmt.Close()

	if _, e := stmt.Exec(params...); e != nil {
		return ErrInserting.Wrap(e)
	}

	return nil
}

func sqlitePartitionedInsert[T any](db *sqliteDB, items []T, buildInsertSQL sqlBuilder[T]) error {
	for _, batch := range partition(items, insertBatchSize) {
		sql, params := buildInsertSQL(batch)

		if e := db.insert(sql, params); e != nil {
			return ErrSQLite.Wrap(e)
		}
	}

	return nil
}

func buildValuesSQL(rowCount, paramCount int) string {
	sb := strings.Builder{}
	params := buildParamsSQL(paramCount)

	for i := 0; i < rowCount; i++ {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(params)
	}

	return sb.String()
}

func buildParamsSQL(paramCount int) string {
	sb := strings.Builder{}
	sb.WriteRune('(')

	for i := 0; i < paramCount; i++ {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune('?')
	}

	sb.WriteRune(')')
	return sb.String()
}

func partition[T any](items []T, batchSize int) [][]T {
	var batches [][]T
	var batch []T

	for _, v := range items {
		if len(batch) == batchSize {
			batches = append(batches, batch)
			batch = nil
		}

		batch = append(batch, v)
	}

	if len(batch) > 0 {
		batches = append(batches, batch)
	}

	return batches
}
