package rotatelogs

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/goravel/file-rotatelogs/v2/strftime"
)

func TestGenerateFn(t *testing.T) {
	// Mock time
	ts := []time.Time{
		{},
		(time.Time{}).Add(24 * time.Hour),
	}

	for _, xt := range ts {
		pattern, err := strftime.New("/path/to/%Y/%m/%d")
		if !assert.NoError(t, err, `strftime.New should succeed`) {
			return
		}
		fn := GenerateFn(pattern, NewClock(xt), 24*time.Hour)
		expected := fmt.Sprintf("/path/to/%04d/%02d/%02d",
			xt.Year(),
			xt.Month(),
			xt.Day(),
		)

		if !assert.Equal(t, expected, fn) {
			return
		}
	}
}

func TestGenerateNowFn(t *testing.T) {
	// Mock time
	ts := []time.Time{
		{},
		(time.Time{}).Add(24 * time.Hour),
		(time.Time{}).Add(25*time.Hour + 55*time.Minute + 15*time.Second),
	}

	for _, xt := range ts {
		pattern, err := strftime.New("/path/to/%Y/%m/%d/%H/%M/%S")
		if !assert.NoError(t, err, `strftime.New should succeed`) {
			return
		}
		fn := GenerateNowFn(pattern, NewClock(xt))
		expected := fmt.Sprintf("/path/to/%04d/%02d/%02d/%02d/%02d/%02d",
			xt.Year(),
			xt.Month(),
			xt.Day(),
			xt.Hour(),
			xt.Minute(),
			xt.Second(),
		)

		if !assert.Equal(t, expected, fn) {
			return
		}
	}
}

func TestCreateFileSuccess(t *testing.T) {
	file, err := CreateFile("testfile.log")
	assert.NoError(t, err)
	assert.NoError(t, file.Close())
	assert.NoError(t, os.Remove("testfile.log"))
}
