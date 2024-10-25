package main

import "database/sql"

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")
	exec(db, "USE cursogo")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER AUTO_INCREMENT, 
		nome VARCHAR(80),
		PRIMARY KEY (id)
	)`)

}
