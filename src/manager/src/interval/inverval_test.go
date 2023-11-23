package interval

import "testing"
import "fmt"

func compareIntervals(a *Interval, b *Interval) bool {
	return a.start == b.start && a.end == b.end && a.step == b.step && a.size == b.size && a.precision == b.precision
}

func testIntervalGeneric(t *testing.T, interval *Interval, expected []Interval, nPartitions uint64) {
	actual := interval.Split(nPartitions)
	if len(actual) != len(expected) {
		t.Errorf("Split(%v) = %v; want %v", nPartitions, actual, expected)
	}
	for i := range actual {
		if !compareIntervals(&actual[i], &expected[i]) {
			t.Errorf("Split(%v) = %v; want %v", nPartitions, &actual[i], &expected[i])
		}
	}
}

func TestIntervalSplit(t *testing.T) {
	tests := []struct {
		name       string
		interval   *Interval
		expected   []Interval
		nPartitions uint64
	}{
		{
			name:       "Test Case 1",
			interval:   NewInterval(0, 10, 1),
			expected:   []Interval{*NewInterval(0, 5, 1), *NewInterval(5, 10, 1)},
			nPartitions: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testIntervalGeneric(t, tt.interval, tt.expected, tt.nPartitions)
		})
	}
}
