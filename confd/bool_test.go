package confd

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeBool(t *testing.T) {
	var b Bool

	b = true
	data, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.Equal(t, "1", string(data[:]))

	b = false
	data, err = json.Marshal(b)
	assert.NoError(t, err)
	assert.Equal(t, "0", string(data[:]))
}

func TestDecodeBool(t *testing.T) {
	var b Bool

	err := json.Unmarshal([]byte("1"), &b)
	assert.NoError(t, err)
	assert.Equal(t, true, bool(b))

	err = json.Unmarshal([]byte("0"), &b)
	assert.NoError(t, err)
	assert.Equal(t, false, bool(b))
}
