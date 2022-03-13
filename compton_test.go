package compton

import (
	"os"
)

var testCompton *Compton
var testDatabase *Database
var testTable *Table
var testCounter int32

func createTestTable(name string) {

	if testTable != nil {
		return
	}

	table, err := testDatabase.CreateTable(name)
	if err != nil {
		panic(err)
	}

	testTable = table
}

func releaseTestTable() {
	testTable.Drop()
	testTable = nil
}

func createTestCompton(name string) {

	if testCompton != nil {
		return
	}

	err := os.RemoveAll("./" + name)
	if err != nil {
		panic(err)
	}

	options := NewOptions()
	options.DatabasePath = "./" + name

	compton, err := NewCompton(options)
	if err != nil {
		panic(err)
	}

	testCompton = compton
}

func releaseTestCompton() {
	releaseTestDatabase()
	testCompton.Close()
	testCompton = nil
}

func createTestDatabase(name string) {

	if testDatabase != nil {
		return
	}

	// Create a new db for benchmark
	db, err := testCompton.CreateDatabase(name)
	if err != nil {
		panic(err)
	}

	testDatabase = db
}

func releaseTestDatabase() {
	releaseTestTable()
	testDatabase.Drop()
	testDatabase = nil
}
