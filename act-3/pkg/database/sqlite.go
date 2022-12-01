package database

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/PaulioRandall/trackable"
)

// TODO: Create SQLite helper library?

var ErrSQLite = trackable.Track("SQLite database error")

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
		return ErrCreating.Wrap(e)
	}

	return nil
}

func (db *sqliteDB) Close() {
	db.conn.Close()
}

func joinLines(lines ...string) string {
	return strings.Join(lines, "\n")
}
