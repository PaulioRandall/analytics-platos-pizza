package sqlite

import (
	"database/sql"

	"github.com/PaulioRandall/analytics-platos-pizza/act-3/pkg/database"
)

func (db *sqliteDB) InsertOrderDetails(orderDetails ...database.OrderDetail) error {
	buildOrderDetailsInsertSQL := func(batch []database.OrderDetail) (sql string, params []any) {
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

	rows, e := db.conn.Query(sql, queryHeadMax)
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
