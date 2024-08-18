package testa

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const statement_create_table = `create table if not exists submissions (
  key text not null primary key,
  value blob
)`

const statement_insert = `insert into submissions (key, value) values (?, ?)`
const statement_retrieve = `select value from submissions where key = (?)`

func db_open(path string, maxCons int) (*sql.DB, error) {
	const options = "?_journal_mode=WAL"
	var err error

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s%s", path, options))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxCons)

	_, err = db.Exec(statement_create_table)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func db_store_entry(db *sql.DB, key string, value []byte) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(statement_insert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(key, value)
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func db_retrieve_key(db *sql.DB, key string) ([]byte, error) {

	stmt, err := db.Prepare(statement_retrieve)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var value []byte
	err = stmt.QueryRow(key).Scan(&value)

	if err != nil {
		return nil, err
	}
	return value, nil
}
