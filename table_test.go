package compton

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
