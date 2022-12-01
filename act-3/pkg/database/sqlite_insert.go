package database

import (
	"strings"
)

type sqlBuilder[T any] func([]T) (sql string, params []any)

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
	return sqlitePartitionedInsert(db, orders, buildOrdersInsertSQL)
}

func buildOrdersInsertSQL(orders []Order) (sql string, params []any) {
	rowCount := len(orders)
	paramCount := 2

	valuesSQL := buildValuesSQL(rowCount, paramCount)
	sql = joinLines(
		`INSERT INTO orders (`,
		`	id,`,
		`	datetime`,
		`) VALUES `+valuesSQL+";",
	)

	for _, v := range orders {
		params = append(params, v.Id, v.Datetime)
	}

	return sql, params
}

func (db *sqliteDB) InsertOrderDetails(orderDetails ...OrderDetail) error {
	return nil
}

func (db *sqliteDB) InsertPizzas(pizzas ...Pizza) error {
	return nil
}

func (db *sqliteDB) InsertPizzaTypes(pizzaTypes ...PizzaType) error {
	return sqlitePartitionedInsert(db, pizzaTypes, buildPizzaTypesInsertSQL)
}

func buildPizzaTypesInsertSQL(pizzaTypes []PizzaType) (sql string, params []any) {
	rowCount := len(pizzaTypes)
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

	for _, v := range pizzaTypes {
		params = append(params, v.Id, v.Name, v.Category, v.Ingredients)
	}

	return sql, params
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
	for _, batch := range partition(items, 256) {
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
