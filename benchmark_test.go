package compton

import (
	"encoding/binary"
	"testing"

	record_type "github.com/BrobridgeOrg/compton/types/record"
	"google.golang.org/protobuf/types/known/structpb"
)

func BenchmarkInternalWrite(b *testing.B) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	key := make([]byte, 4)
	value := []byte("value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		err := testTable.write(key, value)
		if err != nil {
			panic(err)
		}
	}
	testTable.sync()
	b.StopTimer()
}

func BenchmarkInternalRead(b *testing.B) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	key := make([]byte, 4)
	value := []byte("value")

	for i := 0; i < 1000000; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		err := testTable.write(key, value)
		if err != nil {
			panic(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		_, closer, _ := testTable.get(key)
		closer.Close()
	}
	b.StopTimer()
}

func BenchmarkInternalMerge(b *testing.B) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	key := make([]byte, 4)
	value := []byte("value")

	for i := 0; i < 1000000; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		err := testTable.write(key, value)
		if err != nil {
			panic(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		testTable.merge(key, []byte("updated"), func(key []byte, value []byte) []byte {
			return value
		})
	}
	testTable.sync()
	b.StopTimer()
}

func BenchmarkGet(b *testing.B) {

	createTestCompton("test")
	createTestDatabase("test")
	createTestTable("test")
	defer releaseTestCompton()

	key := make([]byte, 4)
	value := []byte("value")

	for i := 0; i < 1000000; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		err := testTable.write(key, value)
		if err != nil {
			panic(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint32(key, uint32(i))
		testTable.Get(key)
	}
	b.StopTimer()
}

func BenchmarkWriteRecord(b *testing.B) {

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

	err := testTable.WriteRecord([]byte("test_key"), r)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := Int64ToBytes(int64(i))

		r.Payload.Map.Fields[0].Value.Value = key
		err := testTable.WriteRecord(key, r)
		if err != nil {
			panic(err)
		}

	}
	testTable.sync()
	b.StopTimer()
}
