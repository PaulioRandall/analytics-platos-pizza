package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/err"
)

var (
	ErrSQLite = err.Track("SQLite database error")
)

type sqliteDB struct {
	conn *sql.DB
}

func OpenSQLiteDatabase(file string) (sqliteDB, error) {
	conn, e := sql.Open("sqlite3", file)
	if e != nil {
		return sqliteDB{}, ErrSQLite.Wrap(e)
	}

	db := sqliteDB{
		conn: conn,
	}

	return db, nil
}

func (db *sqliteDB) InsertMetadata(entry MetadataEntry) error {
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
	return nil, nil
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
