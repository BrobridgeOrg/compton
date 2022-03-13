package compton

import (
	"errors"
	"os"
)

type Compton struct {
	options *Options
	dbPath  string
	dbs     map[string]*Database
}

var (
	ErrDatabaseExistsAlready = errors.New("Database exists already")
	ErrNotFoundDatabase      = errors.New("Not found database")
)

func NewCompton(options *Options) (*Compton, error) {

	// Preparing database directory
	err := os.MkdirAll(options.DatabasePath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	ct := &Compton{
		dbs:     make(map[string]*Database),
		options: options,
	}

	return ct, nil
}

func (ct *Compton) Close() {

	for _, db := range ct.dbs {
		db.Close()
	}

	ct.dbs = make(map[string]*Database)
}

func (ct *Compton) CreateDatabase(dbName string) (*Database, error) {

	if _, ok := ct.dbs[dbName]; ok {
		return nil, ErrDatabaseExistsAlready
	}

	db, err := NewDatabase(ct, dbName)
	if err != nil {
		return nil, err

	}

	ct.dbs[dbName] = db

	return db, nil
}

func (ct *Compton) GetDatabase(dbName string) (*Database, error) {

	if db, ok := ct.dbs[dbName]; ok {
		return db, nil
	}

	return nil, ErrNotFoundDatabase
}

func (ct *Compton) DropDatabase(dbName string) error {

	db, ok := ct.dbs[dbName]
	if !ok {
		return ErrNotFoundDatabase
	}

	err := db.drop()
	if err != nil {
		return err
	}

	delete(ct.dbs, dbName)

	return nil
}
