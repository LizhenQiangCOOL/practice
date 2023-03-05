package go_safety_code

import (
	"database/sql"
	"fmt"
)

// sql注入
// 防止拼接，要使用预编译

func queryById(db *sql.DB, id, pwd string) string {
	var sqlStr = "SELECT Name, Dept From USER WHERE Id=? AND Pwd=?"
	rows, err := db.Query(sqlStr, id, pwd)
	if err != nil {
		fmt.Println("Query USER table error.")
		return ""
	}
	results, _ := rows.Columns()
	return results[0]
}
