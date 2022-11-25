package database

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

// TODO: Create SQLite helper library?
// TODO: Create trackable (error) library

var ErrSQLite = err.Track("SQLite database error")

type sqliteDB struct {
	conn *sql.DB
}

func OpenSQLiteDatabase(file string) (*sqliteDB, error) {
	conn, e := sql.Open("sqlite3", file)
	if e != nil {
		return nil, ErrSQLite.Wrap(e)
	}

	db := &sqliteDB{
		conn: conn,
	}

	if e = db.createTables(); e != nil {
		return nil, ErrSQLite.Wrap(e)
	}

	return db, nil
}

func (db *sqliteDB) createTables() error {
	sql := joinLines(
		`CREATE TABLE metadata (`,
		`	id          INTEGER NOT NULL PRIMARY KEY,`, // Alias for SQLite 'rowid'
		`	table_name  TEXT    NOT NULL,`,
		`	field_name  TEXT    NOT NULL,`,
		`	description TEXT    NOT NULL`,
		`);`,
		``,
		`CREATE TABLE pizza_types (`,
		`	id          TEXT NOT NULL PRIMARY KEY,`,
		`	name        TEXT NOT NULL,`,
		`	category    TEXT NOT NULL,`,
		`	ingredients TEXT NOT NULL`,
		`);`,
		``,
		`CREATE TABLE pizzas (`,
		`	id      TEXT NOT NULL PRIMARY KEY,`,
		`	type_id TEXT NOT NULL,`,
		`	size    TEXT NOT NULL,`,
		`	price   REAL NOT NULL,`,
		`	FOREIGN KEY(type_id) REFERENCES pizza_types(id)`,
		`);`,
		``,
		`CREATE TABLE orders (`,
		`	id       INTEGER NOT NULL PRIMARY KEY,`,
		`	datetime TEXT    NOT NULL`,
		`);`,
		``,
		`CREATE TABLE order_details (`,
		`	id       INTEGER NOT NULL PRIMARY KEY,`,
		`	order_id INTEGER NOT NULL,`,
		`	pizza_id TEXT    NOT NULL,`,
		`	quantity INTEGER NOT NULL,`,
		`	FOREIGN KEY(order_id) REFERENCES orders(id),`,
		`	FOREIGN KEY(pizza_id) REFERENCES pizzas(id)`,
		`);`,
	)

	if _, e := db.conn.Exec(sql); e != nil {
		return ErrCreateOrUpdate.TraceWrap(e, "Could not create tables")
	}

	return nil
}

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
		e = err.Wrap(e, "Failed inserting metadata")
		return ErrSQLite.TrackWrap(ErrInsert, e)
	}

	return nil
}

func (db *sqliteDB) InsertOrders(orders ...Order) error {
	for _, batch := range partition(orders, 256) {
		sql, params := buildBulkOrderInsertSQL(batch)

		if e := db.insert(sql, params); e != nil {
			e = err.Wrap(e, "Failed inserting orders")
			return ErrSQLite.TrackWrap(ErrInsert, e)
		}
	}

	return nil
}

func buildBulkOrderInsertSQL(orders []Order) (sql string, params []any) {
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
	return nil
}

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
		return nil, ErrQuery.Wrap(e)
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
			return nil, ErrResult.Wrap(e)
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
		return nil, ErrQuery.Wrap(e)
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
			return nil, ErrResult.Wrap(e)
		}

		order.Datetime, e = time.Parse(DatetimeFormat, datetimeStr)
		if e != nil {
			return nil, ErrResult.Wrap(e)
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
	return nil, nil
}

func (db *sqliteDB) Close() {
	db.conn.Close()
}

func (db *sqliteDB) insert(sql string, params []any) error {
	stmt, e := db.conn.Prepare(sql)
	if e != nil {
		return ErrInsert.TrackWrap(ErrPrepare, e)
	}
	defer stmt.Close()

	if _, e := stmt.Exec(params...); e != nil {
		return ErrInsert.Wrap(e)
	}

	return nil
}

func joinLines(lines ...string) string {
	return strings.Join(lines, "\n")
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
