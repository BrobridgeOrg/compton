package compton

import "github.com/cockroachdb/pebble"

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
