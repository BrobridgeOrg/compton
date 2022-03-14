package compton

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"

	record_type "github.com/BrobridgeOrg/compton/types/record"
	"github.com/cockroachdb/pebble"
)

var (
	ErrNotFoundRecord = errors.New("Not found record")
	ErrNotFoundEntry  = errors.New("Not found entry")
)

var (
	MetaDataKeyPrefix = []byte("m_")
	RecordKeyPrefix   = []byte("r_")
)

type Table struct {
	Database *Database
	Db       *pebble.DB
	Path     string
	Name     string
	Merge    func([]byte, []byte) []byte

	mergers     sync.Map
	closed      chan struct{}
	isClosed    bool
	isScheduled uint32
	pending     uint32
	timer       *time.Timer
}

func NewTable(db *Database, name string) *Table {
	table := &Table{
		Database: db,
		Name:     name,
		Merge: func(oldValue []byte, newValue []byte) []byte {
			return newValue
		},
		closed:      make(chan struct{}),
		isClosed:    false,
		isScheduled: 0,
		pending:     0,
		timer:       time.NewTimer(time.Second * 10),
	}

	table.timer.Stop()

	return table
}

func (table *Table) syncLoop() {

	table.timer.Reset(time.Second * 10)

	for {

		select {
		case <-table.timer.C:

			if !table.sync() {
				continue
			}

			atomic.StoreUint32(&table.isScheduled, 0)
		case <-table.closed:
			table.timer.Stop()
			close(table.closed)
			return
		}
	}
}

func (table *Table) sync() bool {

	if atomic.LoadUint32(&table.pending) == 0 {
		return false
	}

	table.Db.LogData(nil, pebble.Sync)

	table.timer.Stop()
	table.timer.Reset(time.Second * 10)

	atomic.StoreUint32(&table.pending, 0)

	return true
}

func (table *Table) requestSync() {

	if atomic.LoadUint32(&table.pending) == 0 {
		return
	}

	if atomic.LoadUint32(&table.isScheduled) != 0 {
		return
	}

	atomic.StoreUint32(&table.isScheduled, 1)

	table.timer.Stop()

	// Sync in 0.1 second
	table.timer.Reset(time.Millisecond * 100)
}

func (table *Table) Open() error {

	opts := &pebble.Options{
		Merger: &pebble.Merger{
			Merge: func(key []byte, value []byte) (pebble.ValueMerger, error) {
				m := &Merger{
					table: table,
				}

				fn, ok := table.mergers.LoadAndDelete(BytesToString(key))
				if !ok {
					m.SetHandler(table.Merge)
					return m, m.MergeNewer(value)
				}

				m.SetHandler(fn.(func([]byte, []byte) []byte))

				atomic.AddUint32(&table.pending, 1)

				table.requestSync()

				return m, m.MergeNewer(value)
			},
		},
		//		DisableWAL:    true,
		MaxOpenFiles:  -1,
		LBaseMaxBytes: 512 << 20,
	}

	db, err := pebble.Open(table.Path, opts)
	if err != nil {
		return err
	}

	table.Db = db

	go table.syncLoop()

	return nil
}

func (table *Table) Close() error {

	if table.isClosed {
		return nil
	}

	table.isClosed = true
	table.closed <- struct{}{}
	table.Db.LogData(nil, pebble.Sync)
	return table.Db.Close()
}

func (table *Table) drop() error {

	err := table.Close()
	if err != nil {
		return err
	}

	// Remove files of table
	err = os.RemoveAll(table.Path)
	if err != nil {
		return err
	}

	return nil
}

func (table *Table) Drop() error {
	return table.Database.DropTable(table.Name)
}

func (table *Table) write(key []byte, data []byte) error {

	err := table.Db.Set(key, data, pebble.NoSync)
	if err != nil {
		return err
	}

	atomic.AddUint32(&table.pending, 1)
	table.requestSync()

	return nil
}

func (table *Table) get(key []byte) ([]byte, io.Closer, error) {

	v, closer, err := table.Db.Get(key)
	if err == pebble.ErrNotFound {
		return nil, nil, ErrNotFoundEntry
	}

	return v, closer, err
}

func (table *Table) delete(key []byte) error {
	err := table.Db.Delete(key, pebble.NoSync)
	if err != nil {
		return err
	}

	atomic.AddUint32(&table.pending, 1)

	return nil
}

func (table *Table) list(targetKey []byte) (*Cursor, error) {

	iter := table.Db.NewIter(nil)
	if !iter.SeekGE(targetKey) || !iter.Valid() {
		return nil, ErrNotFoundRecord
	}

	cur := &Cursor{
		iter: iter,
	}

	return cur, nil
}

func (table *Table) merge(key []byte, value []byte, fn func([]byte, []byte) []byte) error {

	table.mergers.Store(BytesToString(key), fn)
	err := table.Db.Merge(key, value, pebble.NoSync)
	if err != nil {
		return err
	}

	return nil
}

func (table *Table) Get(key []byte) ([]byte, error) {

	value, closer, err := table.Db.Get(key)
	if err != nil {
		return nil, err
	}

	data := make([]byte, len(value))
	copy(data, value)

	closer.Close()

	return data, nil
}

func (table *Table) Write(key []byte, data []byte) error {
	return table.write(key, data)
}

func (table *Table) WriteRecord(pkey []byte, r *record_type.Record) error {

	key := append(RecordKeyPrefix, pkey...)

	data, err := record_type.Marshal(r)
	if err != nil {
		return err
	}

	return table.write(key, data)
}

func (table *Table) GetRecord(pkey []byte) (*record_type.Record, error) {

	key := append(RecordKeyPrefix, pkey...)

	value, closer, err := table.get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	r := record_type.Record{}

	err = record_type.Unmarshal(value, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (table *Table) ModifyRecord(pkey []byte, newRecord *record_type.Record) error {

	key := append(RecordKeyPrefix, pkey...)

	err := table.merge(key, []byte(""), func(oldValue []byte, newValue []byte) []byte {

		originRecord := record_type.Record{}

		err := record_type.Unmarshal(oldValue, &originRecord)
		if err != nil {
			fmt.Println(err)
			return oldValue
		}

		return record_type.Merge(&originRecord, newRecord)
	})
	if err != nil {
		return err
	}

	return nil
}

func (table *Table) ListRecords(targetPrimaryKey []byte) (*Cursor, error) {

	targetKey := append(RecordKeyPrefix, targetPrimaryKey...)

	iter := table.Db.NewIter(nil)
	if !iter.SeekGE(targetKey) || !iter.Valid() {
		return nil, ErrNotFoundRecord
	}

	cur := &Cursor{
		iter: iter,
	}

	return cur, nil
}

func (table *Table) SetMetaBytes(key []byte, value []byte) error {
	k := append(MetaDataKeyPrefix, key...)
	return table.write(k, value)
}

func (table *Table) GetMetaBytes(key []byte) ([]byte, error) {

	k := append(MetaDataKeyPrefix, key...)
	value, closer, err := table.get(k)
	if err != nil {
		return nil, err
	}

	data := make([]byte, len(value))
	copy(data, value)

	closer.Close()

	return data, nil
}

func (table *Table) DeleteMeta(key []byte) error {
	k := append(MetaDataKeyPrefix, key...)
	return table.delete(k)
}

func (table *Table) ListMeta(key []byte) (*Cursor, error) {
	k := append(MetaDataKeyPrefix, key...)
	return table.list(k)
}

func (table *Table) SetMetaInt64(key []byte, value int64) error {

	k := append(MetaDataKeyPrefix, key...)

	data := Int64ToBytes(value)

	return table.write(k, data)
}

func (table *Table) GetMetaInt64(key []byte) (int64, error) {

	k := append(MetaDataKeyPrefix, key...)

	value, closer, err := table.get(k)
	if err != nil {
		return 0, err
	}

	data := BytesToInt64(value)

	closer.Close()

	return data, nil
}

func (table *Table) SetMetaUint64(key []byte, value uint64) error {

	k := append(MetaDataKeyPrefix, key...)

	data := Uint64ToBytes(value)

	return table.write(k, data)
}

func (table *Table) GetMetaUint64(key []byte) (uint64, error) {

	k := append(MetaDataKeyPrefix, key...)

	value, closer, err := table.get(k)
	if err != nil {
		return 0, err
	}

	data := BytesToUint64(value)

	closer.Close()

	return data, nil
}

func (table *Table) SetMetaFloat64(key []byte, value float64) error {

	k := append(MetaDataKeyPrefix, key...)

	data := Float64ToBytes(value)

	return table.write(k, data)
}

func (table *Table) GetMetaFloat64(key []byte) (float64, error) {

	k := append(MetaDataKeyPrefix, key...)

	value, closer, err := table.get(k)
	if err != nil {
		return 0, err
	}

	data := BytesToFloat64(value)

	closer.Close()

	return data, nil
}

func (table *Table) SetMetaString(key []byte, value string) error {

	k := append(MetaDataKeyPrefix, key...)

	data := StrToBytes(value)

	return table.write(k, data)
}

func (table *Table) GetMetaString(key []byte) (string, error) {

	k := append(MetaDataKeyPrefix, key...)

	value, closer, err := table.get(k)
	if err != nil {
		return "", err
	}

	data := make([]byte, len(value))
	copy(data, value)

	closer.Close()

	return string(data), nil
}
