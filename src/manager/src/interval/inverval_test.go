package interval

import "testing"

func compareIntervals(a *Interval, b *Interval) bool {
	return a.start == b.start && a.end == b.end && a.step == b.step && a.size == b.size && a.precision == b.precision
}

func testIntervalGeneric(t *testing.T, interval *Interval, expected []Interval, nPartitions uint64) {
	actual := interval.Split(nPartitions)
	if len(actual) != len(expected) {
		t.Errorf("Split(%v) = %v; want %v", nPartitions, actual, expected)
		return
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
			name:       "even split with whole positive numbers", 
			interval:   NewInterval(0, 10, 1),
			expected:   []Interval{*NewInterval(0, 5, 1), *NewInterval(5, 10, 1)},
			nPartitions: 2,
		},
		{
			name:       "even split with whole negative numbers", 
			interval:   NewInterval(-10, 10, 1),
			expected:   []Interval{*NewInterval(-10, 0, 1), *NewInterval(0, 10, 1)},
			nPartitions: 2,
		},
		{
			name:       "even split with whole negative numbers odd split amount",
			interval:   NewInterval(-600, 600, 1),
			expected:   []Interval{*NewInterval(-600, -200, 1), *NewInterval(-200, 200, 1), *NewInterval(200, 600, 1)},
			nPartitions: 3,
		},
		{
			name:       "uneven split", 
			interval:   NewInterval(0, 10, 3),
			expected:   []Interval{*NewInterval(0, 6, 3), *NewInterval(6, 9, 3), *NewInterval(9, 12, 3)},
			nPartitions: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testIntervalGeneric(t, tt.interval, tt.expected, tt.nPartitions)
		})
	}
}
