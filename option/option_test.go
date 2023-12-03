package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionCreation(t *testing.T) {
	opt := New("test", "value")
	assert.NotNil(t, opt)
}

func TestOptionName(t *testing.T) {
	opt := New("test", "value")
	name := opt.Name()
	assert.Equal(t, "test", name)
}

func TestOptionValue(t *testing.T) {
	opt := New("test", "value")
	value := opt.Value()
	assert.Equal(t, "value", value)
}

func TestOptionValueDifferentType(t *testing.T) {
	opt := New("test", 123)
	value := opt.Value()
	assert.Equal(t, 123, value)
}
