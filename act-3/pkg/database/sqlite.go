package database

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

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
		`CREATE TABLE orders (`,
		`	id        INTEGER NOT NULL PRIMARY KEY,`,
		`	datetime  TEXT    NOT NULL`,
		`);`,
	)

	if _, e := db.conn.Exec(sql); e != nil {
		return ErrCreateOrUpdate.TraceWrap(e, "Could not create tables")
	}

	return nil
}

func (db *sqliteDB) InsertMetadata(entries ...MetadataEntry) error {
	sql, params := buildMetadataInsertSQL(entries)

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

func buildMetadataInsertSQL(entries []MetadataEntry) (sql string, params []any) {
	sql = joinLines(
		`INSERT INTO metadata (`,
		`	table_name,`,
		`	field_name,`,
		`	description`,
		`) VALUES`,
	)

	for i, v := range entries {
		if i == 0 {
			sql += " (?, ?, ?)"
		} else {
			sql += ", (?, ?, ?)"
		}

		params = append(params, v.Table, v.Field, v.Description)
	}

	sql += ";"
	return sql, params
}

func (db *sqliteDB) InsertOrders(orders ...Order) error {
	for _, batch := range partition(orders, 256) {
		if e := insertOrderBatch(db, batch); e != nil {
			return ErrInsert.TrackWrap(ErrPrepare, e)
		}
	}

	return nil
}

func insertOrderBatch(db *sqliteDB, batch []Order) error {
	sql, params := buildOrdersInsertSQL(batch)

	stmt, e := db.conn.Prepare(sql)
	if e != nil {
		return err.Wrap(e, "Failed to insert order batch")
	}
	defer stmt.Close()

	if _, e := stmt.Exec(params...); e != nil {
		return ErrInsert.Wrap(e)
	}

	return nil
}

func buildOrdersInsertSQL(orders []Order) (sql string, params []any) {
	sql = joinLines(
		`INSERT INTO orders (`,
		`	id,`,
		`	datetime`,
		`) VALUES`,
	)

	for i, v := range orders {
		if i == 0 {
			sql += " (?, ?)"
		} else {
			sql += ", (?, ?)"
		}

		params = append(params, v.Id, v.Datetime)
	}

	sql += ";"
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

func joinLines(lines ...string) string {
	return strings.Join(lines, "\n")
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
