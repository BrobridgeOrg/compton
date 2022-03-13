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

/*
func (db *Database) getValue(tableName string, key []byte) ([]byte, io.Closer, error) {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return nil, nil, err
	}

	return tableHandle.Db.Get(key)
}

func (db *Database) Delete(tableName string, key []byte) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	return tableHandle.Db.Delete(key, pebble.NoSync)
}

func (db *Database) Put(tableName string, key []byte, data []byte) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	return tableHandle.write(key, data)
}

func (db *Database) PutInt64(tableName string, key []byte, value int64) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	data := Int64ToBytes(value)

	return tableHandle.write(key, data)
}

func (db *Database) GetInt64(tableName string, key []byte) (int64, error) {

	value, closer, err := db.getValue(tableName, key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return 0, nil
		}

		return 0, err
	}

	data := BytesToInt64(value)

	closer.Close()

	return data, nil
}

func (db *Database) PutUint64(tableName string, key []byte, value uint64) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	data := Uint64ToBytes(value)

	return tableHandle.write(key, data)
}

func (db *Database) GetUint64(tableName string, key []byte) (uint64, error) {

	value, closer, err := db.getValue(tableName, key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return 0, nil
		}

		return 0, err
	}

	data := BytesToUint64(value)

	closer.Close()

	return data, nil
}

func (db *Database) PutFloat64(tableName string, key []byte, value float64) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	data := Float64ToBytes(value)

	return tableHandle.write(key, data)
}

func (db *Database) GetFloat64(tableName string, key []byte) (float64, error) {

	value, closer, err := db.getValue(tableName, key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return 0, nil
		}

		return 0, err
	}

	data := BytesToFloat64(value)

	closer.Close()

	return data, nil
}

func (db *Database) GetBytes(tableName string, key []byte) ([]byte, error) {

	value, closer, err := db.getValue(tableName, key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return []byte(""), nil
		}

		return []byte(""), err
	}

	data := make([]byte, len(value))
	copy(data, value)

	closer.Close()

	return data, nil
}

func (db *Database) PutString(tableName string, key []byte, value string) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	data := StrToBytes(value)

	return tableHandle.write(key, data)
}

func (db *Database) GetString(tableName string, key []byte) (string, error) {

	value, closer, err := db.getValue(tableName, key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return "", nil
		}

		return "", err
	}

	data := make([]byte, len(value))
	copy(data, value)

	closer.Close()

	return string(data), nil
}

func (db *Database) List(tableName string, targetKey []byte, callback func(key []byte, value []byte) bool) error {

	tableHandle, err := db.getTable(tableName)
	if err != nil {
		return err
	}

	iter := tableHandle.Db.NewIter(nil)
	for iter.SeekGE(targetKey); iter.Valid(); iter.Next() {
		isContinuous := callback(iter.Key(), iter.Value())

		if !isContinuous {
			break
		}
	}

	return iter.Close()
}
*/
