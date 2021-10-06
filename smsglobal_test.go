package smsglobal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := New("", "")
	assert.Error(t, err)

	s, _ := New("key", "secret")
	assert.Equal(t, "key", s.User.Handler.Key)
	assert.Equal(t, "secret", s.User.Handler.Secret)
}