package holidays

import (
	"reflect"
	"testing"
	"time"
)

func TestEaster(t *testing.T) {
	tests := []struct {
		year uint
		want time.Time
	}{
		{year: 2022, want: time.Date(2022, time.April, 17, 0, 0, 0, 0, time.UTC)},
		{year: 2023, want: time.Date(2023, time.April, 9, 0, 0, 0, 0, time.UTC)},
		{year: 2024, want: time.Date(2024, time.March, 31, 0, 0, 0, 0, time.UTC)},
		{year: 2025, want: time.Date(2025, time.April, 20, 0, 0, 0, 0, time.UTC)},
		{year: 2026, want: time.Date(2026, time.April, 5, 0, 0, 0, 0, time.UTC)},
		{year: 2027, want: time.Date(2027, time.March, 28, 0, 0, 0, 0, time.UTC)},
		{year: 2028, want: time.Date(2028, time.April, 16, 0, 0, 0, 0, time.UTC)},
		{year: 2029, want: time.Date(2029, time.April, 1, 0, 0, 0, 0, time.UTC)},
		{year: 2042, want: time.Date(2042, time.April, 6, 0, 0, 0, 0, time.UTC)},
		{year: 2052, want: time.Date(2052, time.April, 21, 0, 0, 0, 0, time.UTC)},
		{year: 2062, want: time.Date(2062, time.March, 26, 0, 0, 0, 0, time.UTC)},
		{year: 2071, want: time.Date(2071, time.April, 19, 0, 0, 0, 0, time.UTC)},
		{year: 2083, want: time.Date(2083, time.April, 4, 0, 0, 0, 0, time.UTC)},
		{year: 2084, want: time.Date(2084, time.March, 26, 0, 0, 0, 0, time.UTC)},
		{year: 2085, want: time.Date(2085, time.April, 15, 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range tests {
		got := calculateEasterDate(int(tc.year))
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestPrevDayFunc(t *testing.T) {
	tests := []struct {
		Date     time.Time
		Prev     time.Weekday
		expected time.Time
	}{
		{Date: newDate(2022, time.December, 11), Prev: time.Sunday, expected: newDate(2022, time.December, 4)},
		{Date: newDate(2022, time.December, 11), Prev: time.Monday, expected: newDate(2022, time.December, 5)},
		{Date: newDate(2022, time.December, 11), Prev: time.Tuesday, expected: newDate(2022, time.December, 6)},
		{Date: newDate(2022, time.December, 11), Prev: time.Wednesday, expected: newDate(2022, time.December, 7)},
		{Date: newDate(2022, time.December, 11), Prev: time.Thursday, expected: newDate(2022, time.December, 8)},
		{Date: newDate(2022, time.December, 11), Prev: time.Friday, expected: newDate(2022, time.December, 9)},
		{Date: newDate(2022, time.December, 11), Prev: time.Saturday, expected: newDate(2022, time.December, 10)},

		{Date: newDate(2023, time.April, 3), Prev: time.Sunday, expected: newDate(2023, time.April, 2)},
		{Date: newDate(2023, time.April, 3), Prev: time.Saturday, expected: newDate(2023, time.April, 1)},
		{Date: newDate(2023, time.April, 3), Prev: time.Friday, expected: newDate(2023, time.March, 31)},
		{Date: newDate(2023, time.April, 3), Prev: time.Thursday, expected: newDate(2023, time.March, 30)},
		{Date: newDate(2023, time.April, 3), Prev: time.Wednesday, expected: newDate(2023, time.March, 29)},
		{Date: newDate(2023, time.April, 3), Prev: time.Tuesday, expected: newDate(2023, time.March, 28)},
		{Date: newDate(2023, time.April, 3), Prev: time.Monday, expected: newDate(2023, time.March, 27)},
	}
	for _, tc := range tests {
		got := PreviousDayOfWeek(tc.Date, tc.Prev)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
