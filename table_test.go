package compton

import (
	"testing"

	record_type "github.com/BrobridgeOrg/compton/types/record"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestTableLowlevelWrite(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	err := testTable.write([]byte("test_key"), []byte("test_value"))
	if err != nil {
		t.Error(err)
	}

	value, closer, err := testTable.get([]byte("test_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []byte("test_value"), value)

	err = closer.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestTableLowlevelMerge(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	err := testTable.write([]byte("test_key"), []byte("test_value"))
	if err != nil {
		t.Error(err)
	}

	err = testTable.merge([]byte("test_key"), []byte("test_new_value"), func(key []byte, value []byte) []byte {
		return value
	})
	if err != nil {
		t.Error(err)
	}

	value, closer, err := testTable.get([]byte("test_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []byte("test_new_value"), value)

	err = closer.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestTableLowlevelDelete(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	err := testTable.write([]byte("test_key"), []byte("test_value"))
	if err != nil {
		t.Error(err)
	}

	err = testTable.delete([]byte("test_key"))
	if err != nil {
		t.Error(err)
	}

	_, closer, err := testTable.get([]byte("test_key"))
	assert.Equal(t, ErrNotFoundEntry, err)

	if closer != nil {
		err = closer.Close()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestTableList(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	for i := 1; i <= 10; i++ {
		key := Int64ToBytes(int64(i))

		err := testTable.write(key, key)
		if err != nil {
			t.Error(err)
		}
	}

	var counter int64 = 0
	targetKey := Int64ToBytes(int64(1))
	cur, err := testTable.list(targetKey)
	if err != nil {
		t.Error(err)
	}

	for !cur.EOF() {

		counter++

		value := cur.GetData()

		assert.Equal(t, counter, BytesToInt64(value))

		cur.Next()
	}

	assert.Equal(t, int64(10), counter)
}

func TestTableWriteRecord(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	r := record_type.NewRecord()
	meta, _ := structpb.NewStruct(map[string]interface{}{})
	r.Meta = meta
	r.Payload.Map.Fields = []*record_type.Field{
		&record_type.Field{
			Name: "id",
			Value: &record_type.Value{
				Type:  record_type.DataType_STRING,
				Value: []byte("test"),
			},
		},
	}

	err := testTable.WriteRecord([]byte("test_key"), r)
	if err != nil {
		t.Error(err)
	}

	key := append(RecordKeyPrefix, []byte("test_key")...)

	value, closer, err := testTable.get(key)
	if err != nil {
		t.Error(err)
	}

	// Parsing data
	result := record_type.Record{}

	err = record_type.Unmarshal(value, &result)
	if err != nil {
		t.Error(err)
	}

	err = closer.Close()
	if err != nil {
		t.Error(err)
	}

	data, _ := result.GetValueDataByPath("id")
	assert.Equal(t, "test", data.(string))
}

func TestTableGetRecord(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	r := record_type.NewRecord()
	meta, _ := structpb.NewStruct(map[string]interface{}{})
	r.Meta = meta
	r.Payload.Map.Fields = []*record_type.Field{
		&record_type.Field{
			Name: "id",
			Value: &record_type.Value{
				Type:  record_type.DataType_STRING,
				Value: []byte("test"),
			},
		},
	}

	err := testTable.WriteRecord([]byte("test_key"), r)
	if err != nil {
		t.Error(err)
	}

	record, err := testTable.GetRecord([]byte("test_key"))
	if err != nil {
		t.Error(err)
	}

	data, _ := record.GetValueDataByPath("id")

	assert.Equal(t, "test", data.(string))
}

func TestTableModifyRecord(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	r := record_type.NewRecord()
	meta, _ := structpb.NewStruct(map[string]interface{}{})
	r.Meta = meta
	r.Payload.Map.Fields = make([]*record_type.Field, 0)

	r.Payload.Map.Fields = append(r.Payload.Map.Fields,
		&record_type.Field{
			Name: "id",
			Value: &record_type.Value{
				Type:  record_type.DataType_STRING,
				Value: []byte("test"),
			},
		},
		&record_type.Field{
			Name: "name",
			Value: &record_type.Value{
				Type:  record_type.DataType_STRING,
				Value: []byte("name_value"),
			},
		},
	)

	err := testTable.WriteRecord([]byte("test_key"), r)
	if err != nil {
		t.Error(err)
	}

	r.Payload.Map.Fields[1].Value.Value = []byte("modified")
	r.Payload.Map.Fields = append(r.Payload.Map.Fields,
		&record_type.Field{
			Name: "note",
			Value: &record_type.Value{
				Type:  record_type.DataType_STRING,
				Value: []byte("note_value"),
			},
		},
	)
	err = testTable.ModifyRecord([]byte("test_key"), r)
	if err != nil {
		t.Error(err)
	}

	testTable.sync()

	record, err := testTable.GetRecord([]byte("test_key"))
	if err != nil {
		t.Error(err)
	}

	data, _ := record.GetValueDataByPath("id")
	assert.Equal(t, "test", data.(string))

	data, _ = record.GetValueDataByPath("name")
	assert.Equal(t, "modified", data.(string))

	data, _ = record.GetValueDataByPath("note")
	assert.Equal(t, "note_value", data.(string))
}
func TestTableDeleteRecord(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	r := record_type.NewRecord()
	meta, _ := structpb.NewStruct(map[string]interface{}{})
	r.Meta = meta
	r.Payload.Map.Fields = []*record_type.Field{
		&record_type.Field{
			Name: "id",
			Value: &record_type.Value{
				Type:  record_type.DataType_STRING,
				Value: []byte("test"),
			},
		},
	}

	err := testTable.WriteRecord([]byte("test_key"), r)
	if err != nil {
		t.Error(err)
	}

	err = testTable.DeleteRecord([]byte("test_key"))
	if err != nil {
		t.Error(err)
	}

	_, err = testTable.GetRecord([]byte("test_key"))
	assert.Equal(t, ErrNotFoundRecord, err)
}

func TestTableListRecords(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	r := record_type.NewRecord()
	meta, _ := structpb.NewStruct(map[string]interface{}{})
	r.Meta = meta
	r.Payload.Map.Fields = []*record_type.Field{
		&record_type.Field{
			Name: "id",
			Value: &record_type.Value{
				Type:  record_type.DataType_INT64,
				Value: Int64ToBytes(0),
			},
		},
	}

	for i := 1; i <= 10; i++ {
		key := Int64ToBytes(int64(i))

		r.Payload.Map.Fields[0].Value.Value = Int64ToBytes(int64(i))
		err := testTable.WriteRecord(key, r)
		if err != nil {
			t.Error(err)
		}
	}

	var counter int64 = 0
	targetKey := Int64ToBytes(int64(1))
	cur, err := testTable.ListRecords(targetKey)
	if err != nil {
		t.Error(err)
	}

	for !cur.EOF() {

		counter++

		record := cur.GetRecord()
		data, _ := record.GetValueDataByPath("id")

		assert.Equal(t, counter, data.(int64))

		cur.Next()
	}

	assert.Equal(t, int64(10), counter)
}

func TestTableSetMetaBytes(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	if err := testTable.SetMetaBytes([]byte("meta_key"), []byte("meta_value")); err != nil {
		t.Error(err)
	}

	value, err := testTable.GetMetaBytes([]byte("meta_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []byte("meta_value"), value)
}

func TestTableSetMetaInt64(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	if err := testTable.SetMetaInt64([]byte("meta_key"), int64(999999)); err != nil {
		t.Error(err)
	}

	value, err := testTable.GetMetaInt64([]byte("meta_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, int64(999999), value)
}

func TestTableSetMetaUint64(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	if err := testTable.SetMetaUint64([]byte("meta_key"), uint64(999999)); err != nil {
		t.Error(err)
	}

	value, err := testTable.GetMetaUint64([]byte("meta_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, uint64(999999), value)
}

func TestTableSetMetaFloat64(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	if err := testTable.SetMetaFloat64([]byte("meta_key"), float64(999.999)); err != nil {
		t.Error(err)
	}

	value, err := testTable.GetMetaFloat64([]byte("meta_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, float64(999.999), value)
}

func TestTableSetMetaString(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	if err := testTable.SetMetaString([]byte("meta_key"), "test"); err != nil {
		t.Error(err)
	}

	value, err := testTable.GetMetaString([]byte("meta_key"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "test", value)
}

func TestTableDeleteMeta(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	if err := testTable.SetMetaBytes([]byte("meta_key"), []byte("meta_value")); err != nil {
		t.Error(err)
	}

	_, err := testTable.GetMetaBytes([]byte("meta_key"))
	assert.Equal(t, nil, err)

	err = testTable.DeleteMeta([]byte("meta_key"))
	if err != nil {
		t.Error(err)
	}

	_, err = testTable.GetMetaBytes([]byte("meta_key"))
	assert.Equal(t, ErrNotFoundEntry, err)
}

func TestTableListMeta(t *testing.T) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	for i := 1; i <= 10; i++ {
		key := Int64ToBytes(int64(i))

		err := testTable.SetMetaInt64(key, int64(i))
		if err != nil {
			t.Error(err)
		}
	}

	var counter int64 = 0
	targetKey := Int64ToBytes(int64(1))
	cur, err := testTable.ListMeta(targetKey)
	if err != nil {
		t.Error(err)
	}

	for !cur.EOF() {

		counter++

		value := cur.GetData()

		key := Int64ToBytes(int64(counter))
		correctKey := append(MetaDataKeyPrefix, key...)
		assert.Equal(t, correctKey, cur.GetKey())
		assert.Equal(t, counter, BytesToInt64(value))

		cur.Next()
	}

	assert.Equal(t, int64(10), counter)
}
