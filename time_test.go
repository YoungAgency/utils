package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMidnight(t *testing.T) {
	t.Run("Test using time.Now()", func(t *testing.T) {
		now := time.Now()
		midnight := GetMidnight(now)

		// date must be equal to now
		assert.Equal(t, now.Year(), midnight.Year())
		assert.Equal(t, now.Month(), midnight.Month())
		assert.Equal(t, now.Day(), midnight.Day())

		// hours, minutes, seconds and ns offset must be 0
		assert.Equal(t, 0, midnight.Hour())
		assert.Equal(t, 0, midnight.Minute())
		assert.Equal(t, 0, midnight.Second())
		assert.Equal(t, 0, midnight.Nanosecond())
	})
}

func TestGetYesteday(t *testing.T) {
	t.Run("Test using time.Now()", func(t *testing.T) {
		now := time.Now()
		yesterday := GetYesterday(now)

		// date must be equal to now
		assert.Equal(t, now.YearDay()-1, yesterday.YearDay())

		// hours, minutes, seconds and ns offset must be 0
		assert.Equal(t, 0, yesterday.Hour())
		assert.Equal(t, 0, yesterday.Minute())
		assert.Equal(t, 0, yesterday.Second())
		assert.Equal(t, 0, yesterday.Nanosecond())
	})
}
