package compton

import (
	record_type "github.com/BrobridgeOrg/compton/types/record"
	"github.com/cockroachdb/pebble"
)

type Cursor struct {
	iter *pebble.Iterator
}

func (cur *Cursor) EOF() bool {
	return !cur.iter.Valid()
}

func (cur *Cursor) Next() bool {
	return cur.iter.Next()
}

func (cur *Cursor) Close() error {
	return cur.iter.Close()
}

func (cur *Cursor) GetData() []byte {
	return cur.iter.Value()
}

func (cur *Cursor) GetRecord() *record_type.Record {
	value := cur.iter.Value()

	r := record_type.Record{}

	err := record_type.Unmarshal(value, &r)
	if err != nil {
		return nil
	}

	return &r
}
