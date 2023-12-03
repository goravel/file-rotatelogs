package rotatelogs

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRotateLogs(t *testing.T) {
	rl, err := New("test.log", WithMaxAge(24*time.Hour), WithRotationTime(1*time.Hour))
	assert.NoError(t, err)
	assert.NotNil(t, rl)
	assert.NoError(t, rl.Close())
}

func TestWriteLog(t *testing.T) {
	rl, _ := New("test.log", WithMaxAge(24*time.Hour), WithRotationTime(1*time.Hour))
	n, err := rl.Write([]byte("test log"))
	assert.NoError(t, err)
	assert.Equal(t, 8, n)
	assert.NoError(t, rl.Close())
	assert.NoError(t, os.Remove("test.log"))
}

func TestWriteLogWhenRotateFile(t *testing.T) {
	rl, _ := New("test.log", WithMaxAge(24*time.Hour), WithRotationTime(1*time.Hour))
	assert.NoError(t, rl.Rotate())
	n, err := rl.Write([]byte("test log"))
	assert.NoError(t, err)
	assert.Equal(t, 8, n)
	assert.NoError(t, rl.Close())
	assert.NoError(t, os.Remove("test.log"))
}

func TestRotateLogs(t *testing.T) {
	rl, _ := New("test.log", WithMaxAge(24*time.Hour), WithRotationTime(1*time.Hour))
	err := rl.Rotate()
	assert.NoError(t, err)
	assert.NoError(t, rl.Close())
	assert.NoError(t, os.Remove("test.log"))
}

func TestCurrentFileName(t *testing.T) {
	rl, _ := New("test.log", WithMaxAge(24*time.Hour), WithRotationTime(1*time.Hour))
	n, err := rl.Write([]byte("test log"))
	assert.NoError(t, err)
	assert.Equal(t, 8, n)
	filename := rl.CurrentFileName()
	assert.Equal(t, "test.log", filename)
	assert.NoError(t, rl.Close())
	assert.NoError(t, os.Remove("test.log"))
}

func TestClose(t *testing.T) {
	rl, _ := New("test.log", WithMaxAge(24*time.Hour), WithRotationTime(1*time.Hour))
	err := rl.Close()
	assert.NoError(t, err)
}
