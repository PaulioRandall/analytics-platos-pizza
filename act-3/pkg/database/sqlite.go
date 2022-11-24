package database

import (
	"database/sql"
	"strings"

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
	if e := db.createMetadataTable(); e != nil {
		return e
	}

	if e := db.createOrdersTable(); e != nil {
		return e
	}

	if e := db.createOrderDetailsTable(); e != nil {
		return e
	}

	if e := db.createPizzasTable(); e != nil {
		return e
	}

	if e := db.createPizzaTypesTable(); e != nil {
		return e
	}

	return nil
}

func (db *sqliteDB) createMetadataTable() error {
	metadataTable := joinLines(
		`CREATE TABLE metadata (`,
		`	id          INTEGER NOT NULL PRIMARY KEY,`, // Alias for SQLite 'rowid'
		`	table_name  TEXT    NOT NULL,`,
		`	field_name  TEXT    NOT NULL,`,
		`	description TEXT    NOT NULL`,
		`);`,
	)

	if _, e := db.conn.Exec(metadataTable); e != nil {
		return ErrCreateOrUpdate.TraceWrap(e, "Could not create metadata table")
	}

	return nil
}

func (db *sqliteDB) createOrdersTable() error {
	return nil
}

func (db *sqliteDB) createOrderDetailsTable() error {
	return nil
}

func (db *sqliteDB) createPizzasTable() error {
	return nil
}

func (db *sqliteDB) createPizzaTypesTable() error {
	return nil
}

func (db *sqliteDB) InsertMetadata(entry MetadataEntry) error {
	sql := joinLines(
		`INSERT INTO metadata (`,
		`	table_name,`,
		`	field_name,`,
		`	description`,
		`) VALUES (?, ?, ?);`,
	)

	stmt, e := db.conn.Prepare(sql)
	if e != nil {
		return ErrInsert.TrackWrap(ErrPrepare, e)
	}
	defer stmt.Close()

	if _, e := stmt.Exec(entry.Table, entry.Field, entry.Description); e != nil {
		return ErrInsert.Wrap(e)
	}

	return nil
}

func (db *sqliteDB) InsertOrder(order Order) error {
	return nil
}

func (db *sqliteDB) InsertOrderDetail(orderDetail OrderDetail) error {
	return nil
}

func (db *sqliteDB) InsertPizza(pizza Pizza) error {
	return nil
}

func (db *sqliteDB) InsertPizzaType(pizzaType PizzaType) error {
	return nil
}

func (db *sqliteDB) AllMetadata() ([]MetadataEntry, error) {
	sql := joinLines(
		`SELECT`,
		`	table_name,`,
		`	field_name,`,
		`	description`,
		`FROM`,
		`	metadata`,
	)

	rows, e := db.conn.Query(sql)
	if e != nil {
		return nil, ErrQuery.Wrap(e)
	}
	defer rows.Close()

	return scanMetadataRows(rows)
}

func (db *sqliteDB) HeadOrders() ([]Order, error) {
	return nil, nil
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
