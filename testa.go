package testa

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Maximum number of read-only connections permitted
// by read-only handle
var ReadOnlyMaxConns int = 5000

// Read-Only
type Read interface {
	Get(key string) ([]byte, error)
}

// Write-Only
type Write interface {
	Set(key string, data []byte) error
}

// RW Interface
type ReadWrite interface {
	Read
	Write
}

// Database handles
type Db struct {
	readOnly  *sql.DB
	readWrite *sql.DB
}

// Open a database at the given path and recieve
// database handle (should use Close upon completion)
// Use as ReadWrite interface
func Open(path string) (*Db, error) {
	tdb := Db{
		nil,
		nil,
	}

	var err error

	tdb.readWrite, err = db_open(path, 1)
	if err != nil {
		return nil, err
	}

	tdb.readOnly, err = db_open(path, ReadOnlyMaxConns)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &tdb, nil
}

// Close the db
func (tdb *Db) Close() {
	if tdb.readOnly != nil {
		tdb.readOnly.Close()
		tdb.readOnly = nil
	}
	if tdb.readWrite != nil {
		tdb.readWrite.Close()
		tdb.readWrite = nil
	}
}

// Set the key of something
func (tdb *Db) Set(key string, value []byte) error {
	err := db_store_entry(tdb.readWrite, key, value)
	if err != nil {
		return err
	}
	return nil
}

// Set the value of something
func (tdb *Db) Get(key string) ([]byte, error) {
	return db_retrieve_key(tdb.readOnly, key)
}
