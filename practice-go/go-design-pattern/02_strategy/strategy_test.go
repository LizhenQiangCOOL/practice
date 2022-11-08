package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_storage(t *testing.T) {
	data, sensitive := getDataMock()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}

	storage, err := NewStorageStrategy(strategyType)
	assert.NoError(t, err)
	assert.NoError(t, storage.Save("./test.txt", data))
}

func getDataMock() ([]byte, bool) {
	return []byte("test data aaaa"), false
}
