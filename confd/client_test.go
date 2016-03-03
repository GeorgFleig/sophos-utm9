package confd

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializeOptionsAnonymous(t *testing.T) {
	conn := NewAnonymousConn()
	conn.Options.Name = "test"
	bytes, err := json.Marshal(conn.Options)
	assert.NoError(t, err)
	assert.Equal(t, `{"client":"test"}`, string(bytes[:]))
}