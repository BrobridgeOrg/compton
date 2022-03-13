package compton

import (
	"encoding/binary"
	"testing"
)

func BenchmarkWrite(b *testing.B) {

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

func BenchmarkRead(b *testing.B) {

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

func BenchmarkMerge(b *testing.B) {

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
