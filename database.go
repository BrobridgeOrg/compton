package compton

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrNotFoundTable = errors.New("Not found table")
)

type Database struct {
	Name string

	compton *Compton
	options *Options
	dbPath  string
	tables  map[string]*Table
}

func NewDatabase(compton *Compton, dbName string) (*Database, error) {

	db := &Database{
		Name:    dbName,
		compton: compton,
		options: compton.options,
		dbPath:  filepath.Join(compton.options.DatabasePath, dbName),
		tables:  make(map[string]*Table),
	}

	err := db.open()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) open() error {

	err := os.MkdirAll(db.dbPath, os.ModePerm)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(db.dbPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		_, err := db.assertTable(file.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Close() {

	for _, table := range db.tables {
		table.Close()
	}
}

func (db *Database) drop() error {

	// Drop all tables
	for _, table := range db.tables {
		table.drop()
		delete(db.tables, table.Name)
	}

	db.Close()

	// Remove all files
	err := os.RemoveAll(db.dbPath)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Drop() error {
	return db.compton.DropDatabase(db.Name)
}

func (db *Database) DropTable(name string) error {

	table, err := db.getTable(name)
	if err != nil {
		return err
	}

	return table.drop()
}

func (db *Database) assertTable(name string) (*Table, error) {

	table, ok := db.tables[name]
	if !ok {
		return db.CreateTable(name)
	}

	return table, nil
}

func (db *Database) CreateTable(name string) (*Table, error) {

	table := NewTable(db, name)
	table.Path = filepath.Join(db.dbPath, name)
	err := table.Open()
	if err != nil {
		return nil, err
	}

	db.tables[name] = table

	return table, nil
}

func (db *Database) GetTable(name string) (*Table, error) {
	return db.getTable(name)
}

func (db *Database) getTable(name string) (*Table, error) {

	table, ok := db.tables[name]
	if !ok {
		return nil, ErrNotFoundTable
	}

	return table, nil
}
