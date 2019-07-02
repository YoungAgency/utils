package date

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMidnight(t *testing.T) {
	t.Run("Test using time.Now()", func(t *testing.T) {
		now := time.Now().UTC()
		midnight := Midnight(now)

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
		now := time.Now().UTC()
		yesterday := Midnight(Yesterday(now))

		// date must be equal to now
		assert.Equal(t, now.YearDay()-1, yesterday.YearDay())

		// hours, minutes, seconds and ns offset must be 0
		assert.Equal(t, 0, yesterday.Hour())
		assert.Equal(t, 0, yesterday.Minute())
		assert.Equal(t, 0, yesterday.Second())
		assert.Equal(t, 0, yesterday.Nanosecond())
	})
}

func TestWeekEnd(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name  string
		args  args
		wantR time.Time
	}{
		{
			name: "Test weekend",
			args: args{
				t: Time(1562057683000),
			},
			wantR: Time(1562544000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := WeekEnd(tt.args.t); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("WeekEnd() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestWeekStart(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name  string
		args  args
		wantR time.Time
	}{
		{
			name: "Test week start",
			args: args{
				t: Time(1562057683000),
			},
			wantR: Time(1561939200000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := WeekStart(tt.args.t); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("WeekStart() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
