package date

import (
	"reflect"
	"testing"
	"time"
)

func TestGetMidnight(t *testing.T) {
	t.Run("Test using time.Now()", func(t *testing.T) {
		now := time.Now().UTC()
		midnight := Midnight(now)

		// date must be equal to now
		if now.Year() != midnight.Year() {
			t.Errorf("expected year %d, got %d", now.Year(), midnight.Year())
		}
		if now.Month() != midnight.Month() {
			t.Errorf("expected month %d, got %d", now.Month(), midnight.Month())
		}
		if now.Day() != midnight.Day() {
			t.Errorf("expected day %d, got %d", now.Day(), midnight.Day())
		}

		// hours, minutes, seconds and ns offset must be 0
		if midnight.Hour() != 0 {
			t.Errorf("expected hour %d, got %d", 0, midnight.Hour())
		}
		if midnight.Minute() != 0 {
			t.Errorf("expected minute %d, got %d", 0, midnight.Minute())
		}
		if midnight.Second() != 0 {
			t.Errorf("expected second %d, got %d", 0, midnight.Second())
		}
		if midnight.Nanosecond() != 0 {
			t.Errorf("expected nanosecond %d, got %d", 0, midnight.Nanosecond())
		}
	})
}

func TestGetYesteday(t *testing.T) {
	t.Run("Test using time.Now()", func(t *testing.T) {
		now := time.Now().UTC()
		yesterday := Midnight(Yesterday(now))

		// date must be equal to now
		if now.YearDay()-1 != yesterday.YearDay() {
			t.Errorf("expected yearday %d, got %d", now.YearDay()-1, yesterday.YearDay())
		}

		// hours, minutes, seconds and ns offset must be 0
		if yesterday.Hour() != 0 {
			t.Errorf("expected hour %d, got %d", 0, yesterday.Hour())
		}
		if yesterday.Minute() != 0 {
			t.Errorf("expected minute %d, got %d", 0, yesterday.Minute())
		}
		if yesterday.Second() != 0 {
			t.Errorf("expected second %d, got %d", 0, yesterday.Second())
		}
		if yesterday.Nanosecond() != 0 {
			t.Errorf("expected nanosecond %d, got %d", 0, yesterday.Nanosecond())
		}
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
			gotR := WeekEnd(tt.args.t)
			if !reflect.DeepEqual(gotR, tt.wantR) {
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
			gotR := WeekStart(tt.args.t)
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("WeekStart() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
