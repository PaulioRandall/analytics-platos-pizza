package sqlite

import (
	"database/sql"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
)

func (db *sqliteDB) InsertPizzaTypes(pizzaTypes ...database.PizzaType) error {
	buildPizzaTypesInsertSQL := func(batch []database.PizzaType) (sql string, params []any) {
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

	rows, e := db.conn.Query(sql, queryHeadMax)
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
